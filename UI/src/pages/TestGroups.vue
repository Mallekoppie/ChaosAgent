<template>
  <div class="content">
    <md-dialog :md-active.sync="confirmDeleteVisible">
      <md-dialog-title>
        Confirm Delete
      </md-dialog-title>
      <md-dialog-content>
        <div class="md-layout">
          <div class="md-layout-item">
            <h3>Name</h3>
            {{ confirmDeleteName }}
            <h4>Description</h4>
            {{ confirmDeleteDescription }}
          </div>

        </div>
        <div class="md-layout md-alignment-center-right">
          <md-button class="md-danger" @click="testGroupDelete">
            <i class="md-icon md-icon-font md-theme-default">delete</i>
            Confirm
          </md-button>
        </div>

      </md-dialog-content>
    </md-dialog>
    <div class="md-layout">
      <div
          class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-100"
      >
        <md-card>
          <md-card-header data-background-color="green">
            <h4 class="title">Test Groups</h4>
            <p class="category">These tests are grouped together to reduce clutter</p>
          </md-card-header>
          <md-card-content>
            <md-table v-model="testGroups" :table-header-color="green">
              <md-table-row slot="md-table-row" slot-scope="{ item }">
                <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
                <md-table-cell md-label="Description">{{ item.description }}</md-table-cell>
                <md-table-cell md-label="Open">
                  <div class="md-button-content">
                    <md-button @click="openTestGroup(item)">
                      <i class="md-icon md-icon-font md-theme-default">open_in_new</i>
                    </md-button>
                  </div>
                </md-table-cell>
                <md-table-cell md-label="Delete">
                  <md-button class="md-danger" @click="confirmTestGroupDeletion(item)">
                    <div class="md-button-content"><i class="md-icon md-icon-font md-theme-default">close</i></div>
                  </md-button>
                </md-table-cell>
              </md-table-row>
            </md-table>
          </md-card-content>
        </md-card>
      </div>
    </div>

    <!-- Add New Test Group -->
    <md-card v-if="addTestGroupVisible">
      <md-card-header data-background-color="orange">
        <label>New Test Group</label>
      </md-card-header>
      <md-card-content>
        <div class="md-layout-item">
          <md-field>
            <label>Name</label>
            <md-input type="text" v-model="newTestGroupName"></md-input>
          </md-field>
        </div>
        <div class="md-layout-item">
          <md-field>
            <label>Description</label>
            <md-input type="text" v-model="newTestGroupDescription"></md-input>
          </md-field>
        </div>
        <div class="md-layout-item md-size-100 text-right">
          <md-button class="md-raised md-success" @click="newTestGroupSave">Save</md-button>
          <md-button class="md-raised md-danger" @click="addTestGroupVisible = false">Cancel</md-button>
        </div>
        <div class="md-layout-item md-size-100 text-right">

        </div>
      </md-card-content>
    </md-card>

    <div class="md-layout md-alignment-center-right">
      <md-button class="md-primary dropdown-toggle" @click="addTestGroupVisible = true">
        <i class="md-icon md-icon-font md-theme-default">add</i>
        Add Test Group
      </md-button>
    </div>
  </div>
</template>

<script>
import {dataStore} from "../shared/datastoretemp";
import {shared} from "@/main.js"

const uuid = require("uuid");

export default {
  name: "TestGroups",
  data() {
    return {
      testGroups: [],
      addTestGroupVisible: false,
      newTestGroupName: "",
      newTestGroupDescription: "",
      confirmDeleteVisible: false,
      confirmDeleteName: "",
      confirmDeleteDescription: "",
      confirmDeleteId: ""
    };
  },
  created() {
    this.testGroups = dataStore.getTestGroups();
  },
  methods: {
    openTestGroup(input) {
      shared.router.push({name: 'TestCollectionList', params: input})
    },
    newTestGroupSave() {
      let newTestGroup = dataStore.createEmptyTestGroup();
      newTestGroup.id = uuid();
      newTestGroup.name = this.newTestGroupName;
      newTestGroup.description = this.newTestGroupDescription;

      this.testGroups.push(newTestGroup);
      this.addTestGroupVisible = false;
      this.newTestGroupName = "";
      this.newTestGroupDescription = "";
    },
    testGroupDelete() {
      let index = this.testGroups.findIndex(t => t.id == this.confirmDeleteId);
      this.testGroups.splice(index, 1);
      this.confirmDeleteVisible = false
      this.confirmDeleteName = "";
      this.confirmDeleteDescription = "";
      this.confirmDeleteId = "";
      //  TODO: Call API to delete
    },
    confirmTestGroupDeletion(test) {
      this.confirmDeleteName = test.name;
      this.confirmDeleteDescription = test.description;
      this.confirmDeleteId = test.id;
      this.confirmDeleteVisible = true;
    }
  }
}
</script>

<style scoped>

</style>