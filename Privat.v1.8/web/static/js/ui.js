window.closeAllMenus = function (exceptMenu = null) {
    document.querySelectorAll('.dropdown-menu.show').forEach(menu => {
        if (menu !== exceptMenu) {
            menu.classList.remove('show');
        }
    });
};

window.getCurrentCipherKey = function (selectorBtn) {
    return selectorBtn.dataset.cipherKey || 'scytale';
};

window.getCurrentOperation = function () {
    return document.getElementById('cipher-op')?.value || 'encrypt';
};

window.getCurrentWorkMode = function () {
    return document.body.dataset.workMode || 'text' ;
}

window.updateActionButton = function (actionBtn) {
    const op = window.getCurrentOperation();
    actionBtn.textContent = op === 'decrypt' ? 'Расшифровать' : 'Зашифровать';
};

window.updateCardanoVisibility = function (selectorBtn) {
    const isCardano = window.getCurrentCipherKey(selectorBtn) === 'cardano';
    const isDecrypt = window.getCurrentOperation() === 'decrypt';
    const codeGroup = document.getElementById('cardano-code-group');

    if (!codeGroup) {
        return;
    }

    if (isCardano && isDecrypt) {
        codeGroup.classList.remove('is-hidden');
    } else {
        codeGroup.classList.add('is-hidden');
    }
};

window.updateUIState = function (selectorBtn, actionBtn) {
    window.updateActionButton(actionBtn);
    window.updateCardanoVisibility(selectorBtn);
};

window.initInnerDropdowns = function (settingsArea, selectorBtn, actionBtn) {
    const dropdowns = settingsArea.querySelectorAll('[data-dropdown]');

    dropdowns.forEach(dropdown => {
        const trigger = dropdown.querySelector('.dropdown-trigger');
        const menu = dropdown.querySelector('.dropdown-menu');
        const items = dropdown.querySelectorAll('.dropdown-item');

        if (!trigger || !menu) {
            return;
        }

        trigger.addEventListener('click', (e) => {
            e.stopPropagation();
            window.closeAllMenus(menu);
            menu.classList.toggle('show');
        });

        items.forEach(item => {
            item.addEventListener('click', () => {
                const value = item.dataset.value || '';
                const text = item.textContent.trim();
                const targetId = trigger.dataset.target;

                trigger.textContent = text;
                trigger.dataset.value = value;

                if (targetId) {
                    const hiddenInput = document.getElementById(targetId);
                    if (hiddenInput) {
                        hiddenInput.value = value;
                    }
                }

                menu.classList.remove('show');
                window.updateUIState(selectorBtn, actionBtn);
            });
        });
    });

    window.updateUIState(selectorBtn, actionBtn);
};

window.setCipher = function (cipherKey, buttonText, selectorBtn, settingsArea, actionBtn) {
    selectorBtn.textContent = buttonText;
    selectorBtn.dataset.apiName = window.apiNames[cipherKey] || '';
    selectorBtn.dataset.cipherKey = cipherKey;
    settingsArea.innerHTML = window.configs[cipherKey] || '';

    window.initInnerDropdowns(settingsArea, selectorBtn, actionBtn);
    window.updateUIState(selectorBtn, actionBtn);
};