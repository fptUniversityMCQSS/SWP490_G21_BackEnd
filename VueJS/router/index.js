import Vue from 'vue'
import Router from 'vue-router'
import CompHome from "../components/CompHomeL";
import CompHomeL from "../components/CompHomeL";
import CompAPI from "../components/CompAPI";
import CompQA from "../components/CompQA";
import CompListKnow from "../components/CompListKnow";
import CompUpload from "../components/CompUpload";
import CompHistory from "../components/CompHistory";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/home', component: CompHome
    },
    {
      path: '/homelogin', component: CompHomeL
    },
    {
      path: "/q&a", component: CompQA
    },
    {
      path: "/api", component: CompAPI
    },
    {
      path: "/listknow", component: CompListKnow
    },
    {
      path: "/uploadknow", component: CompUpload
    },
    {
      path: "/history", component: CompHistory
    }
  ]
})
