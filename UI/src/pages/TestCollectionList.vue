<template>
  <div class="content">
    <div class="md-layout">
      <div
          class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-100"
      >
        <md-card>
          <md-card-header data-background-color="green">
            <h4 class="title">{{ testGroupName }}</h4>
            <p class="category">{{ testGroupDescription }}</p>
          </md-card-header>
          <md-card-content>
            <md-table v-model="testCollections" :table-header-color="green">
              <md-table-row slot="md-table-row" slot-scope="{ item }">
                <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
                <md-table-cell md-label="Description">{{ item.description }}</md-table-cell>
                <md-table-cell md-label="Open">
                  <div class="md-button-content">
                    <md-button @click="openTestCollection(item)">
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
import {dataStore} from "@/shared/datastoretemp";
import {shared} from "@/main";

export default {
  name: "TestCollectionList",
  props:{
    testGroupInput: Object
  },
  data(){
    return {
      testCollections:[],
      testGroupName:'',
      testGroupDescription:''
    }
  },
  created() {
    this.testGroupName = this.testGroupInput.name;
    this.testGroupDescription = this.testGroupInput.description;
    this.testCollections = dataStore.getTestCollections()
  },
  methods:{
    openTestCollection(input){
      shared.router.push({name:'TestCollection'})
    }
  }
}
</script>