<template>
  <div class="content">
    <md-card class="md-card-plain">
      <md-card-header data-background-color="green">
        <h4 class="md-title">{{ testCollection.name }}</h4>
        <p class="category">
          {{ testCollection.description }}
        </p>

      </md-card-header>
    </md-card>
    <div class="md-layout md-alignment-center-right">
        <md-button class="md-primary" @click="addTest">
          <i class="md-icon md-icon-font md-theme-default">add</i>
          Add test
        </md-button>
    </div>


    <md-card v-for="test in testCollection.tests" v-bind:key="test.id">
      <md-card-header data-background-color="blue">
        <h4 class="title">{{ test.name }}</h4>
        <p class="category">{{ test.description }}</p>
      </md-card-header>

      <md-card-content>
        <div class="md-layout">
          <div class="md-layout-item md-small-size-100 md-size-33">
            <md-field>
              <label>Url</label>
              <md-input v-model="test.url"></md-input>
            </md-field>
          </div>
          <div class="md-layout-item md-small-size-100 md-size-33">
            <md-field class="select">
              <label>Method</label>
              <md-select v-model="test.method">
                <md-option value="GET">GET</md-option>
                <md-option value="PUT">PUT</md-option>
                <md-option value="POST">POST</md-option>
                <md-option value="PATCH">PATCH</md-option>
                <md-option value="DELETE">DELETE</md-option>
              </md-select>
            </md-field>


          </div>
          <div class="md-layout-item md-small-size-100 md-size-33">
            <md-field>
              <label>Response Code</label>
              <md-input v-model="test.responseCode" type="text"></md-input>
            </md-field>
          </div>
          <div class="md-layout-item md-small-size-100 md-size-100">
            <md-table v-model="test.headers" :table-header-color="green">
              <md-table-toolbar>
                <h1 class="md-title">Headers</h1>
                <md-button class="md-primary">
                  <i class="md-icon md-icon-font md-theme-default">add</i>
                  ADD
                </md-button>
              </md-table-toolbar>
              <md-table-row slot="md-table-row" slot-scope="{ item }">
                <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
                <md-table-cell md-label="Description">{{ item.value }}</md-table-cell>
                <md-table-cell md-label="Delete">
                  <div class="md-button-content">
                    <md-button class="md-danger">
                      <i class="md-icon md-icon-font md-theme-default">close</i>
                    </md-button>
                  </div>
                </md-table-cell>
              </md-table-row>
            </md-table>
          </div>
          <div class="md-layout-item md-size-100">
            <md-field maxlength="5">
              <label>Request Body</label>
              <md-textarea v-model="test.body"></md-textarea>
            </md-field>
          </div>
          <div class="md-layout-item md-size-100">
            <md-field maxlength="5">
              <label>Response Body</label>
              <md-textarea v-model="test.responseBody"></md-textarea>
            </md-field>
          </div>
          <div class="md-layout-item md-size-100 text-right">
            <md-button class="md-raised md-danger">Delete</md-button>
          </div>
        </div>
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import {dataStore} from "@/shared/datastoretemp";

export default {
  name: "TestCollection",
  props: {
    testCollectionIdInput: Object
  },
  data() {
    return {
      testCollection: Object
    }
  },
  created() {
    this.testCollection = dataStore.getTestDetail()
  },
  methods:{
    addTest(){
      this.testCollection.tests.push( {
        id: 6,
        name: "test 1",
        description:"some description of the test",
        method: "GET",
        url: "http://localhost:9000/bla",
        body: "",
        headers: [
          {
            id: "1",
            name: "Authorization",
            value: "Basic asdlifjnaklsdnjf"
          },
          {
            id: "2",
            name: "Content-Type",
            value: "application/json"
          }
        ],
        responseCode: 200,
        responseBody: ""
      })
    }
  }
}
</script>