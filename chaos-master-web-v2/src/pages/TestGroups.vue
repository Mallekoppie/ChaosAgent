<template>

  <card>
    <template slot="header">
      <div class="row">
        <h6 class="title d-inline col-md-1">Test Groups</h6>
        <p class="card-category d-inline col-md-10">These are just used to group various test collections</p>
        <base-dropdown menu-on-right=""
                       tag="div"
                       title-classes="btn btn-link btn-icon"
                       aria-label="Settings menu"
                       class="col-md-1 align-right">
          <i slot="title" class="tim-icons icon-settings-gear-63"></i>
          <a class="dropdown-item" href="#/add-testgroup">Add new Test Group</a>
        </base-dropdown>
      </div>
    </template>
    <div class="table-full-width table-responsive">
      <base-table :data="testGroups"
                  thead-classes="text-primary">
        <template slot-scope="{row}">
          <td>
            <p class="title">{{ row.name }}</p>
            <p class="text-muted">{{ row.description }}</p>
          </td>
          <td class="td-actions text-right">
            <base-button type="link" aria-label="edit button" @click="viewTestCollections(row)">
              <i class="tim-icons icon-bullet-list-67"></i>
            </base-button>
            <base-button type="link" aria-label="edit button" @click="editTestGroup(row)">
              <i class="tim-icons icon-pencil"></i>
            </base-button>
            <base-button type="link" aria-label="edit button" @click="deleteTestGroup(row)">
              <i class="tim-icons icon-alert-circle-exc"></i>
            </base-button>
            <base-button type="danger" v-if="row.deleteVisible" @click="confirmDeleteTestGroup(row)">
              Press to confirm deletion
            </base-button>
          </td>
        </template>
      </base-table>
    </div>
  </card>

</template>
<script>
import {data} from "@/shared/datastore.js";
import {BaseTable} from '@/components'
import router from "@/router";

export default {
  props: {
    agentInput: {}
  },
  components: {
    BaseTable
  },
  data() {
    return {
      testGroups: [],
      nameState: null,
      newName: "",
      newDescription: ""
    }
  },
  async created() {
    // this.testGroups = await data.getAllTestGroups();
    await this.loadTestGroups();
  },
  computed: {},
  methods: {
    async loadTestGroups() {
      this.testGroups = [];
      this.testGroups = await data.getAllTestGroups();
    },
    deleteTestGroup(testGroup) {
      testGroup.deleteVisible = true;
    },
    async confirmDeleteTestGroup(testGroup){
      console.log(testGroup)
      try {
        await data.deleteTestGroup(testGroup.id);

        let index = this.testGroups.findIndex(h => h.id == testGroup.id);
        this.testGroups.splice(index, 1);
      } catch (e) {
        console.log(e)
      }
    },
    editTestGroup(testGroup) {
      router.push({name: 'add-testgroup', params: {testgroupInput: testGroup}})
    },
    viewTestCollections(testGroup) {
      router.push({name: 'testcollections', params: {testGroupInput: testGroup}})
    },
  },
  mounted() {

  },
  beforeDestroy() {

  }
};
</script>
<style>
</style>
