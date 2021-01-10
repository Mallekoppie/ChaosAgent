<template>

  <div class="row">
    <div class="col">
      <card type="tasks">
        <template slot="header">
          <h6 class="title d-inline">Test Collections</h6>
          <p class="card-category d-inline">Test Collections in this Group</p>
          <base-dropdown menu-on-right=""
                         tag="div"
                         title-classes="btn btn-link btn-icon"
                         aria-label="Settings menu">
            <i slot="title" class="tim-icons icon-settings-gear-63"></i>
            <a class="dropdown-item" @click="addNewTestCollection">Add new Test Collection</a>
          </base-dropdown>
        </template>
        <div class="table-full-width table-responsive">
          <base-table :data="testGroup.testCollections"
                      thead-classes="text-primary">
            <template slot-scope="{row}">
              <td>
                <p class="title">{{ row.name }}</p>

              </td>
              <td>
                <p class="text-muted">{{ row.description }}</p>
              </td>
              <td class="td-actions text-right">
                <base-button type="link" aria-label="edit button" @click="editTestCollection(row)">
                  <i class="tim-icons icon-pencil"></i>
                </base-button>
                <base-button type="link" aria-label="edit button" @click="deleteTestCollection(row)">
                  <i class="tim-icons icon-alert-circle-exc"></i>
                </base-button>
                <base-button type="danger" v-if="row.deleteVisible" @click="confirmDeleteTestCollection(row)">
                  Press to confirm deletion
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
  props: {
    testGroupInput: {}
  },
  components: {
    BaseTable
  },
  data() {
    return {
      testGroup: {}
    }
  },
  async created() {
    await this.loadTestGroup();
  },
  computed: {},
  methods: {
    async loadTestGroup() {
      try {
        let result = await data.getTestGroup(this.testGroupInput.id)
        console.log(result)
        this.testGroup = result;
      } catch (e) {
        this.showErrors = true;
        this.errorMessage = e;
        return
      }

    },
    deleteTestCollection(testCollection) {
      testCollection.deleteVisible = true;
    },
    async confirmDeleteTestCollection(testCollection) {
      await data.deleteTestCollection(testCollection.id)

      let index = this.testGroup.testCollections.findIndex(h => h.id == testCollection.id);
      this.testGroup.testCollections.splice(index, 1);
    },
    addNewTestCollection() {
      router.push({name: 'edit-testcollection', params: {testGroupIdInput: this.testGroup.id}})
    },
    editTestCollection(testCollection) {
      console.log(testCollection)
      router.push({name: 'edit-testcollection', params: {testCollectionInput: testCollection}})
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
