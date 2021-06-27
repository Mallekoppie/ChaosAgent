import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import TestGroups from "../views/TestGroups";
import TestCollections from "../views/TestCollections";
import TestCollection from "@/views/TestCollection";

const parseTestCollectionsProps = r => ({
  testGroupInput: r.params
});

const parseTestCollectionProps = r => ({
  testCollectionInputId: r.params
});

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/test-groups',
    name: 'TestGroups',
    component: TestGroups
  },
  {
    path: '/test-collections',
    name: 'TestCollections',
    component: TestCollections,
    props: parseTestCollectionsProps
  },
  {
    path: '/test-collection',
    name: 'TestCollection',
    component: TestCollection,
    props: parseTestCollectionProps
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
