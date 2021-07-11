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
                <md-button class="md-primary" @click="testHeaderAdd(test)">
                  <i class="md-icon md-icon-font md-theme-default">add</i>
                  ADD
                </md-button>
              </md-table-toolbar>
              <md-table-row slot="md-table-row" slot-scope="{ item }">
                <md-table-cell md-label="Name">
                  <md-field>
                    <label>Header Name</label>
                    <md-input type="text" v-model="item.name" class="md-input"/>
                  </md-field>
                </md-table-cell>
                <md-table-cell md-label="Value">
                  <md-field>
                    <label>Header Value</label>
                    <md-input type="text" v-model="item.value" class="md-input"/>
                  </md-field>
                </md-table-cell>
                <md-table-cell md-label="Delete">
                  <div class="md-button-content">
                    <md-button class="md-danger" @click="testHeaderDelete(test, item)">
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
            <md-button class="md-raised md-danger" @click="testDelete(test)">Delete</md-button>
          </div>
        </div>
      </md-card-content>
    </md-card>

    <md-card v-if="addTestVisible">
      <md-card-header data-background-color="orange">
        <label>New Test</label>
      </md-card-header>
      <md-card-content>
        <div class="md-layout-item">
          <md-field>
            <label>Name</label>
            <md-input type="text" v-model="newTestName"></md-input>
          </md-field>
        </div>
        <div class="md-layout-item">
          <md-field>
            <label>Description</label>
            <md-input type="text" v-model="newTestDescription"></md-input>
          </md-field>
        </div>
        <div class="md-layout-item md-size-100 text-right">
          <md-button class="md-raised md-success" @click="newTestSave">Save</md-button>
          <md-button class="md-raised md-danger" @click="newTestCancel">Cancel</md-button>
        </div>
        <div class="md-layout-item md-size-100 text-right">

        </div>
      </md-card-content>
    </md-card>

    <div class="md-layout md-alignment-center-right">
      <md-button class="md-primary dropdown-toggle" @click="showNewTestCard">
        <i class="md-icon md-icon-font md-theme-default">add</i>
        Add test
      </md-button>
    </div>

    <md-card class="md-layout md-alignment-center-right">
      <md-card-header data-background-color="green">
        <label>Save or Delete</label>
      </md-card-header>
      <md-card-content>
        <div class="md-layout md-alignment-center-right">
          <md-button class="md-success">
            Save
          </md-button>
          <md-button class="md-danger">
            Cancel
          </md-button>
        </div>
        <div class="md-layout-item">

        </div>
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import {dataStore} from "@/shared/datastoretemp";

const uuid = require("uuid");

export default {
  name: "TestCollection",
  props: {
    testCollectionIdInput: Object
  },
  data() {
    return {
      testCollection: Object,
      addTestVisible: false,
      newTestName: '',
      newTestDescription: ''
    }
  },
  created() {
    this.testCollection = dataStore.getTestDetail();
  },
  methods: {
    showNewTestCard() {
      this.addTestVisible = true;

    },
    newTestSave() {
      let newTest = dataStore.createEmptyTest();
      newTest.id = uuid();
      newTest.name = this.newTestName;
      newTest.description = this.newTestDescription;
      this.testCollection.tests.push(newTest);
      this.newTestName = '';
      this.newTestDescription = '';
      this.addTestVisible = false;
    },
    newTestCancel() {
      this.newTestName = '';
      this.newTestDescription = '';
      this.addTestVisible = false;
    },
    testHeaderAdd(test) {
      if (test.headers == null) {
        test.headers = [];
      }

      test.headers.push({
        id: uuid(),
        name: "",
        value: ""
      });
    },
    testHeaderDelete(test, header) {
      let index = test.headers.findIndex(h => h.id == header.id);
      test.headers.splice(index, 1);
    },
    testDelete(test) {
      let index = this.testCollection.tests.findIndex(t => t.id == test.id);
      this.testCollection.tests.splice(index, 1);
    }
  }
}
</script>