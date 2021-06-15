// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'


import './assets/css/bootstrap.css'
import './assets/vendors/linericon/style.css'
import './assets/css/font-awesome.min.css'
import './assets/vendors/owl-carousel/owl.carousel.min.css'
import './assets/vendors/lightbox/simpleLightbox.css'
import './assets/vendors/nice-select/css/nice-select.css'
import './assets/vendors/animate-css/animate.css'
import './assets/vendors/jquery-ui/jquery-ui.css'
import './assets/css/style.css'
import './assets/css/responsive.css'
//
import './assets/js/jquery-3.2.1.min.js'
import './assets/js/popper.js'
import './assets/js/stellar.js'
import './assets/vendors/nice-select/js/jquery.nice-select.min.js'
import './assets/vendors/owl-carousel/owl.carousel.min.js'
import './assets/js/jquery.ajaxchimp.min.js'
import './assets/js/mail-script.js'
import './assets/vendors/jquery-ui/jquery-ui.js'
import './assets/vendors/counter-up/jquery.waypoints.min.js'
import './assets/vendors/counter-up/jquery.counterup.js'

Vue.config.productionTip = false


/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: {App},
  template: '<App/>'
})
