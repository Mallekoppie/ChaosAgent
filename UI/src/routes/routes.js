import DashboardLayout from "@/pages/Layout/DashboardLayout.vue";

import Dashboard from "@/pages/Dashboard.vue";
import TestGroups from "../pages/TestGroups";
import TestCollectionList from "@/pages/TestCollectionList";
import TestCollection from "@/pages/TestCollection";

const parseTestCollectionsProps = r => ({
  testGroupInput: r.params
});

const routes = [
  {
    path: "/",
    component: DashboardLayout,
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        name: "Dashboard",
        component: Dashboard
      },
      {
        path: "testgroups",
        name: "TestGroups",
        component: TestGroups
      },
      {
        path: "testcollectionlist",
        name: "TestCollectionList",
        component: TestCollectionList,
        props: parseTestCollectionsProps
      },
      {
        path: "testcollection",
        name: "TestCollection",
        component: TestCollection
      }
    ]
  }
];

export default routes;
