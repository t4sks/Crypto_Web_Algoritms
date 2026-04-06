document.addEventListener('DOMContentLoaded', () => {
    const selectorBtn = document.getElementById('current-cipher');
    const cipherMenu = document.getElementById('cipher-menu');
    const settingsArea = document.getElementById('dynamic-setting');
    const actionBtn = document.getElementById('encrypt-btn');

    selectorBtn.addEventListener('click', (e) => {
        e.stopPropagation();
        window.closeAllMenus(cipherMenu);
        cipherMenu.classList.toggle('show');
    });

    cipherMenu.querySelectorAll('.dropdown-item').forEach(item => {
        item.addEventListener('click', () => {
            const value = item.dataset.value;
            const text = item.textContent.trim();

            window.setCipher(value, text, selectorBtn, settingsArea, actionBtn);
            cipherMenu.classList.remove('show');
        });
    });

    window.addEventListener('click', () => {
        window.closeAllMenus();
    });

    window.setCipher('scytale', 'Scytale', selectorBtn, settingsArea, actionBtn);

    actionBtn.addEventListener('click', async() =>{
        const mode = window.getCurrentWorkMode();

        if (mode === 'file') {
            await window.handlerFileExecution();
            return;
        }

        await window.handleExecution();

    });
});