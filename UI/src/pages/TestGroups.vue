<template>
  <div class="content">
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
                  <div class="md-button-content"><i class="md-icon md-icon-font md-theme-default">close</i></div>
                </md-table-cell>
              </md-table-row>
            </md-table>
          </md-card-content>
        </md-card>
      </div>
    </div>
  </div>
</template>

<script>
import {dataStore} from "../shared/datastoretemp";
import {shared} from "@/main.js"

export default {
  name: "TestGroups",
  data() {
    return {
      testGroups: []
    };
  },
  created() {
    this.testGroups = dataStore.getTestGroups();
  },
  methods: {
    openTestGroup(input){
      shared.router.push({name:'TestCollectionList', params:input})
    }
  }
}
</script>

<style scoped>

</style>