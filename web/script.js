(function() {
    'use strict';

    // ---- state ----
    let stations = [];
    let currentStationId = null;
    let audioElement = null;
    let isPlaying = false;
    let currentPage = 0;
    const LIMIT = 12;
    let totalStations = 0;

    // DOM refs
    const grid = document.getElementById('stationGrid');
    const searchInput = document.getElementById('searchInput');
    const genreFilter = document.getElementById('genreFilter');
    const countryFilter = document.getElementById('countryFilter');
    const searchBtn = document.getElementById('searchBtn');
    const prevPageBtn = document.getElementById('prevPageBtn');
    const nextPageBtn = document.getElementById('nextPageBtn');
    const pageInfo = document.getElementById('pageInfo');
    const prevPageItem = document.getElementById('prevPageItem');
    const nextPageItem = document.getElementById('nextPageItem');
    const currentStationName = document.getElementById('currentStationName');
    const playPauseBtn = document.getElementById('playPauseBtn');
    const stopBtn = document.getElementById('stopBtn');
    const volumeSlider = document.getElementById('volumeSlider');
    const submitStationBtn = document.getElementById('submitStationBtn');
    const stationNameInput = document.getElementById('stationNameInput');
    const stationUrlInput = document.getElementById('stationUrlInput');
    const stationGenreInput = document.getElementById('stationGenreInput');
    const stationCountryInput = document.getElementById('stationCountryInput');
    const toastEl = document.getElementById('liveToast');
    const toastMessage = document.getElementById('toastMessage');
    const toast = new bootstrap.Toast(toastEl, { delay: 3000 });

    // ---- helpers ----
    function generateFakeStations(count = 22) {
    const genres = ['rock', 'jazz', 'pop', 'classical', 'electronic', 'talk', 'indie', 'ambient'];
    const countries = ['USA', 'UK', 'Russia', 'France', 'Germany', 'Canada', 'Australia', 'Brazil', 'Japan'];
    const names = ['Radio Nova', 'FIP', 'BBC 6 Music', 'KEXP', 'NTS', 'SomaFM', 'Jazz FM', 'Classic FM', 'Radio Paradise', 'Triple J', 'KCRW', 'WFMU', 'Rinse FM', 'Dublab', 'Lofi Radio', 'Pulse', 'Wave', 'Metro', 'Cosmic', 'Ambient'];
    const list = [];
    for (let i = 0; i < count; i++) {
    const name = names[i % names.length] + (i > 10 ? ' ' + (i+1) : '');
    list.push({
    id: `station-${Date.now()}-${i}-${Math.random().toString(36).substring(2,6)}`,
    name: name,
    url: `https://example.com/stream${i}.mp3`,
    genre: genres[i % genres.length],
    country: countries[i % countries.length],
    added_at: new Date(Date.now() - i * 86400000).toISOString(),
});
}
    return list;
}

    function loadStations() {
    const stored = localStorage.getItem('radioStations');
    if (stored) {
    try { stations = JSON.parse(stored); } catch(e) { stations = generateFakeStations(22); }
} else {
    stations = generateFakeStations(22);
}
    totalStations = stations.length;
    renderPage(0);
}

    function persistStations() {
    localStorage.setItem('radioStations', JSON.stringify(stations));
    totalStations = stations.length;
}

    function getFilteredStations() {
    const search = searchInput.value.toLowerCase().trim();
    const genre = genreFilter.value.toLowerCase().trim();
    const country = countryFilter.value.toLowerCase().trim();
    return stations.filter(st => {
    let match = true;
    if (search) match = match && st.name.toLowerCase().includes(search);
    if (genre) match = match && (st.genre || '').toLowerCase().includes(genre);
    if (country) match = match && (st.country || '').toLowerCase().includes(country);
    return match;
});
}

    function escapeHtml(text) {
    if (!text) return '';
    return String(text).replace(/[&<>"]/g, function(m) {
    if (m === '&') return '&amp;'; if (m === '<') return '&lt;'; if (m === '>') return '&gt;'; if (m === '"') return '&quot;';
    return m;
});
}

    // ---- render ----
    function renderPage(page) {
    const filtered = getFilteredStations();
    totalStations = filtered.length;
    const start = page * LIMIT;
    const end = Math.min(start + LIMIT, totalStations);
    const pageItems = filtered.slice(start, end);
    currentPage = page;

    grid.innerHTML = '';
    if (pageItems.length === 0) {
    grid.innerHTML = `<div class="text-center text-secondary py-5" style="grid-column:1/-1;">No stations found</div>`;
} else {
    pageItems.forEach(st => {
    const card = document.createElement('div');
    card.className = `station-card ${currentStationId === st.id ? 'active' : ''}`;
    card.dataset.id = st.id;
    card.innerHTML = `
            <div class="artwork"><i class="bi bi-radio"></i></div>
            <div class="station-name">${escapeHtml(st.name)}</div>
            <span class="genre-tag">${escapeHtml(st.genre || 'mixed')}</span>
            <div class="country-flag"><i class="bi bi-geo-alt"></i> ${escapeHtml(st.country || 'world')}</div>
            <div class="play-indicator">${currentStationId === st.id && isPlaying ? '🔊 playing' : ''}</div>
          `;
    card.addEventListener('click', () => playStation(st.id));
    grid.appendChild(card);
});
}

    // pagination
    const totalPages = Math.ceil(totalStations / LIMIT) || 1;
    pageInfo.textContent = `${page+1} / ${totalPages}`;
    prevPageItem.classList.toggle('disabled', page === 0);
    nextPageItem.classList.toggle('disabled', page >= totalPages - 1);
}

    // ---- player ----
    function playStation(stationId) {
    const station = stations.find(s => s.id === stationId);
    if (!station) return;

    if (currentStationId === stationId && audioElement && isPlaying) {
    pauseAudio();
    return;
}

    if (audioElement) {
    audioElement.pause();
    audioElement.src = '';
    audioElement = null;
    isPlaying = false;
}

    const audio = new Audio(station.url);
    audio.volume = parseFloat(volumeSlider.value);
    audio.crossOrigin = 'anonymous';
    audio.addEventListener('error', () => {
    showToast('Could not load stream. Check the URL.');
    resetPlayerUI();
});
    audio.addEventListener('ended', () => {
    // don't auto-replay
});

    audioElement = audio;
    currentStationId = station.id;
    currentStationName.textContent = station.name;
    playPauseBtn.innerHTML = '<i class="bi bi-pause-fill"></i>';
    isPlaying = true;
    audio.play().catch(() => {});
    renderPage(currentPage);
}

    function pauseAudio() {
    if (audioElement) {
    audioElement.pause();
    isPlaying = false;
    playPauseBtn.innerHTML = '<i class="bi bi-play-fill"></i>';
    renderPage(currentPage);
}
}

    function stopAudio() {
    if (audioElement) {
    audioElement.pause();
    audioElement.src = '';
    audioElement = null;
    isPlaying = false;
    currentStationId = null;
    currentStationName.textContent = 'Select a station';
    playPauseBtn.innerHTML = '<i class="bi bi-play-fill"></i>';
    renderPage(currentPage);
}
}

    function resetPlayerUI() {
    if (audioElement) {
    audioElement.pause();
    audioElement.src = '';
    audioElement = null;
}
    isPlaying = false;
    currentStationId = null;
    currentStationName.textContent = 'Select a station';
    playPauseBtn.innerHTML = '<i class="bi bi-play-fill"></i>';
    renderPage(currentPage);
}

    function togglePlayPause() {
    if (!audioElement || !currentStationId) {
    const filtered = getFilteredStations();
    if (filtered.length) playStation(filtered[0].id);
    return;
}
    if (isPlaying) {
    pauseAudio();
} else {
    if (audioElement) {
    audioElement.play().catch(() => {});
    isPlaying = true;
    playPauseBtn.innerHTML = '<i class="bi bi-pause-fill"></i>';
    renderPage(currentPage);
} else {
    const station = stations.find(s => s.id === currentStationId);
    if (station) playStation(station.id);
}
}
}

    // ---- add station ----
    function addStation(name, url, genre, country) {
    if (!name.trim() || !url.trim()) {
    showToast('Name and URL are required.');
    return false;
}
    if (stations.some(s => s.url === url.trim())) {
    showToast('Station with this URL already exists.');
    return false;
}
    const newStation = {
    id: `station-${Date.now()}-${Math.random().toString(36).substring(2,8)}`,
    name: name.trim(),
    url: url.trim(),
    genre: genre.trim() || 'unknown',
    country: country.trim() || 'world',
    added_at: new Date().toISOString(),
};
    stations.push(newStation);
    persistStations();
    renderPage(currentPage);
    showToast(`Station "${newStation.name}" added!`);
    return true;
}

    function showToast(msg) {
    toastMessage.textContent = msg;
    toast.show();
}

    // ---- event binding ----
    searchBtn.addEventListener('click', () => renderPage(0));
    genreFilter.addEventListener('change', () => renderPage(0));
    countryFilter.addEventListener('input', () => { /* wait for search */ });
    searchInput.addEventListener('keyup', (e) => { if (e.key === 'Enter') renderPage(0); });

    prevPageBtn.addEventListener('click', (e) => {
    e.preventDefault();
    if (currentPage > 0) renderPage(currentPage - 1);
});
    nextPageBtn.addEventListener('click', (e) => {
    e.preventDefault();
    const totalPages = Math.ceil(getFilteredStations().length / LIMIT);
    if (currentPage < totalPages - 1) renderPage(currentPage + 1);
});

    playPauseBtn.addEventListener('click', togglePlayPause);
    stopBtn.addEventListener('click', stopAudio);

    volumeSlider.addEventListener('input', function() {
    if (audioElement) audioElement.volume = parseFloat(this.value);
});

    submitStationBtn.addEventListener('click', () => {
    const name = stationNameInput.value;
    const url = stationUrlInput.value;
    const genre = stationGenreInput.value;
    const country = stationCountryInput.value;
    const ok = addStation(name, url, genre, country);
    if (ok) {
    const modal = bootstrap.Modal.getInstance(document.getElementById('addStationModal'));
    modal.hide();
    stationNameInput.value = '';
    stationUrlInput.value = '';
    stationGenreInput.value = '';
    stationCountryInput.value = '';
}
});

    // keyboard shortcut: space
    document.addEventListener('keydown', (e) => {
    if (e.target.tagName === 'INPUT' || e.target.tagName === 'SELECT') return;
    if (e.code === 'Space') { e.preventDefault(); togglePlayPause(); }
});

    // ---- init ----
    loadStations();

    // handle audio ended to reset UI if needed
    // (we rely on pause/stop)
})();