import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Agents from "../views/Agents.vue";
import Tests from "../views/Tests.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home
  },
  {
    path: "/agents",
    name: "agents",
    component: Agents
  },
  {
    path: "/tests",
    name: "tests",
    component: Tests
  },
  {
    path: "/about",
    name: "about",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
