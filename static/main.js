$(document).ready(function() {
    const animationContainer = $('.animation-container');
    const animationCompleted = sessionStorage.getItem('animationCompleted');

    if (!animationCompleted && animationContainer.length) {
        animationContainer.find('.intro').addClass('go');

        // Заблокировать скролл
        let scrollLocked = true;
        $('body').css('overflow', 'hidden');

        // Запретить скроллинг колесиком мыши
        $(window).on('mousewheel DOMMouseScroll', function(e) {
            if (scrollLocked) {
                e.preventDefault();
            }
        });

        // Здесь используйте setTimeout, чтобы убрать анимацию после определенного времени
        setTimeout(function() {
            // Разблокировать скролл и разрешить скроллинг колесиком мыши
            scrollLocked = false;
            $('body').css('overflow', 'auto');
            $(window).off('mousewheel DOMMouseScroll');

            animationContainer.fadeOut(500, function() {
                // После завершения анимации разблокировать скролл
                scrollLocked = false;
                $('body').css('overflow', 'auto');
                sessionStorage.setItem('animationCompleted', 'true'); // Установить флаг анимации в sessionStorage
                animationContainer.hide(); // Скрыть элемент, но не удалять его
            });
        }, 4800); // Установите здесь время в миллисекундах, через которое анимация должна уйти
    }
});

// Обработка перехода на другие страницы
$(window).on('beforeunload', function() {
    sessionStorage.removeItem('animationCompleted'); // Удалить флаг анимации из sessionStorage
    // Не удаляйте элемент здесь, так как он будет удален в setTimeout в предыдущем блоке кода
});