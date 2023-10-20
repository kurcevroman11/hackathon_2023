let scrollLocked = false;

$(document).ready(function() {
    $('.intro').addClass('go');

    // Заблокировать скролл
    scrollLocked = true;
    $('body').css('overflow', 'hidden');

    // Запретить скроллинг колесиком мыши
    $(window).on('mousewheel DOMMouseScroll', function (e) {
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

        $('.animation-container').fadeOut(500, function() {
            // После завершения анимации разблокировать скролл
            scrollLocked = false;
            $('body').css('overflow', 'auto');
        });
    }, 4800); // Установите здесь время в миллисекундах, через которое анимация должна уйти
});