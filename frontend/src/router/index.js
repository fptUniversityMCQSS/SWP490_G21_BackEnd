import Vue from 'vue'
import Router from 'vue-router'
import CompLogin from "../components/CompLogin";
import CompRegister from "../components/CompRegister";
import CompIndex from "../components/CompIndex";

Vue.use(Router)

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/',
      component: CompIndex
    },
    {
      path: '/login',
      component: CompLogin
    },
    {
      path: '/register',
      component: CompRegister
    }
  ]
})
