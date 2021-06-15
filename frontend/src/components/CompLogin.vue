<template>
  <div>
    <comp-header/>

    <!--================Home Banner Area =================-->
    <section class="banner_area">
      <div class="banner_inner d-flex align-items-center">
        <div class="container">
          <div class="banner_content text-center">
            <h2>Login/Register</h2>
            <div class="page_link">
              <a href="index.html">Home</a>
              <a href="login.html">Login</a>
            </div>
          </div>
        </div>
      </div>
    </section>
    <!--================End Home Banner Area =================-->

    <!--================Login Box Area =================-->
    <section class="login_box_area p_120">
      <div class="container">
        <div class="row">
          <div class="col-lg-6">
            <div class="login_box_img">
              <img class="img-fluid" src="../assets/img/login.jpg" alt="">
              <div class="hover">
                <h4>New to our website?</h4>
                <p>There are advances being made in science and technology everyday, and a good example of this is
                  the</p>
                <router-link to="/register" class="main_btn">Create an Account</router-link>
              </div>
            </div>
          </div>
          <div class="col-lg-6">
            <div class="login_form_inner">
              <h3>Log in to enter</h3>


              <form class="row login_form" @submit="loginData" method="post">
                <div class="col-md-12 form-group">
                  <input type="text" class="form-control" name="username" placeholder="Username"
                         v-model="posts.username">
                </div>
                <div class="col-md-12 form-group">
                  <input type="text" class="form-control" name="password" placeholder="Password"
                         v-model="posts.password">
                </div>
                <div class="col-md-12 form-group">
                  <div class="creat_account">
                    <input type="checkbox" id="f-option2" name="selector">
                    <label for="f-option2">Keep me logged in</label>
                  </div>
                </div>
                <div class="col-md-12 form-group">
                  <button type="submit" value="submit" class="btn submit_btn">Log In</button>
                  <a href="#">Forgot Password?</a>
                </div>
              </form>


            </div>
          </div>
        </div>
      </div>
    </section>
    <!--================End Login Box Area =================-->

    <!--================ Subscription Area ================-->
    <section class="subscription-area section_gap">
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-lg-8">
            <div class="section-title text-center">
              <h2>Subscribe for Our Newsletter</h2>
              <span>We won’t send any kind of spam</span>
            </div>
          </div>
        </div>
        <div class="row justify-content-center">
          <div class="col-lg-6">
            <div>
              <form target="_blank" novalidate
                    action="https://spondonit.us12.list-manage.com/subscribe/post?u=1462626880ade1ac87bd9c93a&id=92a4423d01"
                    method="get" class="subscription relative">
                <input type="email" name="EMAIL" placeholder="Email address" onfocus="this.placeholder = ''"
                       onblur="this.placeholder = 'Email address'"
                       required="">
                <!-- <div style="position: absolute; left: -5000px;">
                    <input type="text" name="b_36c4fd991d266f23781ded980_aefe40901a" tabindex="-1" value="">
                  </div> -->
                <button type="submit" class="newsl-btn">Get Started</button>
                <div class="info"></div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </section>
    <!--================ End Subscription Area ================-->

    <comp-footer/>

  </div>
</template>

<script>
import CompHeader from "./CompHeader";
import CompFooter from "./CompFooter";
import axios from 'axios'
import VueAxios from 'vue-axios'
import Vue from "vue";

Vue.use(VueAxios, axios)

export default {
  name: "CompLogin",
  components: {CompFooter, CompHeader},
  data() {
    return {
      posts: {
        username: null,
        password: null
      }
    }
  },

  methods: {
    loginData(e) {
      // axios({
      //   method: 'post',
      //   url: 'http://localhost:1323/login',
      //   data: {
      //     username: 'binh',
      //     password: '123'
      //   }
      // })
      //   .then((result) => {
      //     console.warn(result)
      //   })
      //   .catch((er) => {
      //     console.log(er)
      //   })

      var axios = require('axios');
      var FormData = require('form-data');
      var data = new FormData();
      data.append('username', 'binh');
      data.append('password', '123');

      var config = {
        method: 'post',
        url: 'http://localhost:1323/login',
        headers: {
          ...data.getHeaders()
        },
        data : data
      };

      axios(config)
        .then(function (response) {
          console.log(JSON.stringify(response.data));
        })
        .catch(function (error) {
          console.log(error);
        });
      e.preventDefault();
    }
  }
}
</script>

<style scoped>

</style>
