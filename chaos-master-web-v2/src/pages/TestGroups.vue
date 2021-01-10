<template>

  <div class="row">
    <div class="col">
      <card type="tasks" :header-classes="{'text-right': isRTL}">
        <template slot="header">
          <h6 class="title d-inline">Test Groups</h6>
          <p class="card-category d-inline">These are just used to group various test collections</p>
          <base-dropdown menu-on-right=""
                         tag="div"
                         title-classes="btn btn-link btn-icon"
                         aria-label="Settings menu"
                         :class="{'float-left': isRTL}">
            <i slot="title" class="tim-icons icon-settings-gear-63"></i>
            <a class="dropdown-item" href="#/add-testgroup">Add new Test Group</a>
            <a class="dropdown-item" href="#pablo">Delete all Test Groups</a>
          </base-dropdown>
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
                <base-button type="link" aria-label="edit button" @click="editTestGroup(row)">
                  <i class="tim-icons icon-pencil"></i>
                </base-button>
              </td>
            </template>
          </base-table>
        </div>
      </card>
    </div>
  </div>

</template>
<script>
import {data} from "@/shared/datastore.js";
import {BaseTable} from '@/components'
import router from "@/router";

export default {
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
  computed: {
    enableRTL() {
      return this.$route.query.enableRTL;
    },
    isRTL() {
      return this.$rtl.isRTL;
    }
  },
  methods: {
    async loadTestGroups() {
      this.testGroups = [];
      this.testGroups = await data.getAllTestGroups();
    },
    async deleteTestGroup(id) {
      await data.deleteTestGroup(id);

      let index = this.testGroups.findIndex(h => h.id == id);
      this.testGroups.splice(index, 1);
    },
    editTestGroup(testGroup) {
      router.push({name: 'add-testgroup', params: {testgroupInput: testGroup}})
    }
  },
  mounted() {

  },
  beforeDestroy() {

  }
};
</script>
<style>
</style>
