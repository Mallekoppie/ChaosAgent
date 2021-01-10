import DashboardLayout from "@/layout/dashboard/DashboardLayout.vue";
// GeneralViews
import NotFound from "@/pages/NotFoundPage.vue";

// Admin pages
const Dashboard = () => import(/* webpackChunkName: "dashboard" */"@/pages/Dashboard.vue");
const TestGroup = () => import(/* webpackChunkName: "dashboard" */"@/pages/TestGroups.vue");
const AddTestGroup = () => import(/* webpackChunkName: "dashboard" */"@/pages/AddTestGroup.vue");
const Agents = () => import(/* webpackChunkName: "dashboard" */"@/pages/Agents.vue");
const AddAgent = () => import(/* webpackChunkName: "dashboard" */"@/pages/AddAgent");
const Profile = () => import(/* webpackChunkName: "common" */ "@/pages/Profile.vue");
const Notifications = () => import(/* webpackChunkName: "common" */"@/pages/Notifications.vue");
const Icons = () => import(/* webpackChunkName: "common" */ "@/pages/Icons.vue");
const Maps = () => import(/* webpackChunkName: "common" */ "@/pages/Maps.vue");
const Typography = () => import(/* webpackChunkName: "common" */ "@/pages/Typography.vue");
const TableList = () => import(/* webpackChunkName: "common" */ "@/pages/TableList.vue");

const parseAgentProps = r => ({
  agentInput: r.params.agentInput
});

const parseTestGroupProps = r => ({
  testgroupInput: r.params.testgroupInput
});

const routes = [
  {
    path: "/",
    component: DashboardLayout,
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        name: "dashboard",
        component: Dashboard
      },
      {
        path: "testgroups",
        name: "testgroups",
        component: TestGroup
      },
      {
        path: "add-testgroup",
        name: "add-testgroup",
        component: AddTestGroup,
        props: parseTestGroupProps
      },
      {
        path: "agents",
        name: "agents",
        component: Agents
      },
      {
        path: "add-agent",
        name: "add-agent",
        component: AddAgent,
        props: parseAgentProps
      },
      {
        path: "profile",
        name: "profile",
        component: Profile
      },
      {
        path: "notifications",
        name: "notifications",
        component: Notifications
      },
      {
        path: "icons",
        name: "icons",
        component: Icons
      },
      {
        path: "maps",
        name: "maps",
        component: Maps
      },
      {
        path: "typography",
        name: "typography",
        component: Typography
      },
      {
        path: "table-list",
        name: "table-list",
        component: TableList
      }
    ]
  },
  { path: "*", component: NotFound },
];

/**
 * Asynchronously load view (Webpack Lazy loading compatible)
 * The specified component must be inside the Views folder
 * @param  {string} name  the filename (basename) of the view to load.
function view(name) {
   var res= require('../components/Dashboard/Views/' + name + '.vue');
   return res;
};**/

export default routes;
