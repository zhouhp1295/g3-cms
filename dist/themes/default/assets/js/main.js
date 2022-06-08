if( !window.console ){
    window.console = {
        log: function(){}
    }
}



/* 
 * jsui
 * ====================================================
*/
jsui.bd = $('body')
jsui.is_signin = jsui.bd.hasClass('logged-in') ? true : false;
jsui.roll = [3, 4, 5]

if( $('.widget-nav').length ){
    $('.widget-nav li').each(function(e){
        $(this).hover(function(){
            $(this).addClass('active').siblings().removeClass('active')
            $('.widget-navcontent .item:eq('+e+')').addClass('active').siblings().removeClass('active')
        })
    })
}


/* 
 * rollbar
 * ====================================================
*/
jsui.rb_comment = ''
if (jsui.bd.hasClass('comment-open')) {
    jsui.rb_comment = "<li><a href=\"javascript:(scrollTo('#comments',-15));\"><i class=\"fa fa-comments\"></i></a><h6>去评论<i></i></h6></li>"
}

jsui.bd.append('\
    <div class="m-mask"></div>\
    <div class="rollbar"><ul>'
    +jsui.rb_comment+
    '<li><a href="javascript:(scrollTo());"><i class="fa fa-angle-up"></i></a><h6>去顶部<i></i></h6></li>\
    </ul></div>\
')

var scroller = $('.rollbar')
$(window).scroll(function() {
    document.documentElement.scrollTop + document.body.scrollTop > 200 ? scroller.fadeIn() : scroller.fadeOut();
})


/* 
 * bootstrap
 * ====================================================
*/


require.config({  
    paths : {  
        bootstrap : '/themes/assets/js/bootstrap.min'
    } 
});


require(['bootstrap'], function(bootstrap) {
    $('.user-welcome').tooltip({
        container: 'body',
        placement: 'bottom'
    })
})



/* 
 * search
 * ====================================================
*/
$('.search-show').bind('click', function(){
    var sbox = $('.site-search')
    $(this).parent().toggleClass('active')
    $(this).find('.fa').toggleClass('fa-remove')
    sbox.toggleClass('active')
    if( sbox.hasClass('active') ){
        sbox.find('input').focus()
    }
})

/* 
 * single
 * ====================================================
*/
if (jsui.bd.hasClass('single')) {
    require(['bootstrap'], function(bootstrap) {
        var _sidebar = $('.sidebar')
        if (_sidebar) {
            var h1 = 15,
                h2 = 30
            var rollFirst = _sidebar.find('.widget:eq(' + (jsui.roll[0] - 1) + ')')
            var sheight = rollFirst.height()

            rollFirst.on('affix-top.bs.affix', function() {
                rollFirst.css({
                    top: 0
                })
                sheight = rollFirst.height()

                for (var i = 1; i < jsui.roll.length; i++) {
                    var item = jsui.roll[i] - 1
                    var current = _sidebar.find('.widget:eq(' + item + ')')
                    current.removeClass('affix').css({
                        top: 0
                    })
                };
            })

            rollFirst.on('affix.bs.affix', function() {
                rollFirst.css({
                    top: h1
                })

                for (var i = 1; i < jsui.roll.length; i++) {
                    var item = jsui.roll[i] - 1
                    var current = _sidebar.find('.widget:eq(' + item + ')')
                    current.addClass('affix').css({
                        top: sheight + h2
                    })
                    sheight += current.height() + 15
                };
            })

            rollFirst.affix({
                offset: {
                    top: _sidebar.height(),
                    bottom: $('.footer').outerHeight()
                }
            })


        }
    })
}


$('.plinks a').each(function(){
    var imgSrc = $(this).attr('href')+'/favicon.ico'
    $(this).prepend( '<img src="'+imgSrc+'">' )
})


/* 
 * page nav
 * ====================================================
*/
if( jsui.bd.hasClass('page-template-pagesnav-php') ){

    $('#navs .items a').attr('target', '_blank')

    require(['bootstrap'], function(bootstrap) {
        $('#navs nav ul').affix({
            offset: {
                top: $('#navs nav ul').offset().top,
                bottom: $('.footer').height() + $('.footer').css('padding-top').split('px')[0]*2
            }
        })
    })

    if( location.hash ){
        var index = location.hash.split('#')[1]
        $('#navs nav .item-'+index).addClass('active')
        scrollTo( '#navs .items .item-'+index )
    }
    $('#navs nav a').each(function(e){
        $(this).click(function(){
            scrollTo( '#navs .items .item-'+$(this).parent().index() )
            $(this).parent().addClass('active').siblings().removeClass('active')
        })
    })
}


/* 
 * page search
 * ====================================================
*/
if( jsui.bd.hasClass('search-results') ){
    var val = $('.searchform .search-input').val()
    var reg = eval('/'+val+'/i')
    $('.excerpt h2 a, .excerpt .note').each(function(){
        $(this).html( $(this).text().replace(reg, function(w){ return '<b>'+w+'</b>' }) )
    })
}


/* 
 * phone
 * ====================================================
*/
$('.m-icon-nav').on('click', function(){
    jsui.bd.toggleClass('m-nav-show')
})

$('.m-mask').on('click', function(){
    jsui.bd.removeClass('m-nav-show')
})


/* functions
 * ====================================================
 */

function scrollTo(name, add, speed) {
    if (!speed) speed = 300
    if (!name) {
        $('html,body').animate({
            scrollTop: 0
        }, speed)
    } else {
        if ($(name).length > 0) {
            $('html,body').animate({
                scrollTop: $(name).offset().top + (add || 0)
            }, speed)
        }
    }
}


