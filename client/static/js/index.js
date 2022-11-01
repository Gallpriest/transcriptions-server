function initSearch() {
    const searchBtn = document.querySelector('.js-search-btn');
    const searchInput = document.querySelector('.js-search-input');

    async function searchListener() {
        const { value } = searchInput;
        
        if (value.trim() === '') return;

        const body = JSON.stringify({ value });
        try {
            const res = await fetch('http://localhost:4040/search', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Content-Length': body.length,
                },
                body,
            });
            const json = await res.json();
            console.log(json)
        } catch (e) {
            console.error(e)
        }
    };

    searchBtn.addEventListener('click', searchListener);
}

function init() {
    initSearch();
}

document.addEventListener('DOMContentLoaded', init);