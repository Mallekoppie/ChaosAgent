import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Agents from "../views/Agents.vue";
import TestGroups from "../views/TestGroups.vue";
import Tests from "../views/Tests.vue";
import TestDetail from "../views/TestDetail.vue";

Vue.use(VueRouter);

const parseTestCollectionProps = r => ({
  testCollectionInput: r.params.testCollectionInput,
  testGroupId: r.params.testGroupId
});

const parseTestGroupProps = r => ({
  testGroupInput: r.params.testGroupInput
});

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
    path: "/testgroups",
    name: "testgroups",
    component: TestGroups
  },
  {
    path: "/tests",
    name: "tests",
    component: Tests,
    props: parseTestGroupProps
  },
  {
    path: "/testDetail",
    name: "test-detail",
    component: TestDetail,
    props: parseTestCollectionProps
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
