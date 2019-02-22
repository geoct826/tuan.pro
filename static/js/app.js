$(document).foundation();

// Patch for a Bug in v6.3.1
$(window).on('changed.zf.mediaquery', function() {
  $('.is-dropdown-submenu.invisible').removeClass('invisible');
});

/*
var homepageExist = !!document.getElementById("homepage");

if (homepageExist) {
var reqPad = 420 - document.getElementById("homepage").clientHeight;
if (reqPad > 0 ) {
		console.log("padding required");
		document.getElementById("callout-background").style.marginTop = reqPad + "px"; 
};
};
*/

$(function() {
  $('a[href*="#"]:not([href="#"])').click(function() {
    if (location.pathname.replace(/^\//,'') == this.pathname.replace(/^\//,'') && location.hostname == this.hostname) {
      var target = $(this.hash);
      target = target.length ? target : $('[name=' + this.hash.slice(1) +']');
      if (target.length) {
        $('html, body').animate({
          scrollTop: target.offset().top
        }, 1000);
        return false;
      }
    }
  });
});

$('html, body').animate({ scrollTop: 0 }, 'fast');