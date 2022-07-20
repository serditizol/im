$(document).ready(function(){

    $('.top_menu_block').slick({
        infinite: true,
        slidesToShow: 9,
        slidesToScroll: 9
    });

    $('.main_slider_block').slick({
        infinite: false,
        slidesToShow: 1,
        slidesToScroll: 1,
        dots: true,
    });

});