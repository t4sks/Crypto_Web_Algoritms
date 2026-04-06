document.addEventListener('DOMContentLoaded', () => {
    const modeSwitch = document.getElementById('mode-switch');
    const textBtn = document.getElementById('mode-text-btn');
    const fileBtn = document.getElementById('mode-file-btn');

    const textPanel = document.getElementById('text-panel');
    const filePanel = document.getElementById('file-panel');

    const actionBtn = document.getElementById('encrypt-btn');

    if (!modeSwitch || !textBtn || !fileBtn || !textPanel || !filePanel || !actionBtn) {
        return;
    }

    function setMode(mode) {
        const isFile = mode === 'file';

        modeSwitch.dataset.mode = mode;

        textBtn.classList.toggle('active', !isFile);
        fileBtn.classList.toggle('active', isFile);

        textPanel.classList.toggle('is-hidden', isFile);
        filePanel.classList.toggle('is-hidden', !isFile);

        actionBtn.textContent = isFile ? 'Обработать файл' : 'Зашифровать';

        document.body.dataset.workMode = mode;
    }

    textBtn.addEventListener('click', () => setMode('text'));
    fileBtn.addEventListener('click', () => setMode('file'));

    setMode('text');
});