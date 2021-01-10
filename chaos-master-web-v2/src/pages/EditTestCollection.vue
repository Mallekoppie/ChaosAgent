<template>
  <div>


    <card>
      <h5 slot="header" class="title">Edit TestCollection</h5>
      <div class="row">
        <div class="col-md-5 pr-md-1">
          <base-input label="ID (generated)"
                      placeholder="id"
                      v-model="testCollection.id"
                      disabled>
          </base-input>
        </div>
        <div class="col-md-5 pr-md-1">
          <base-input label="Group ID (generated)"
                      placeholder="id"
                      v-model="testCollection.groupId"
                      disabled>
          </base-input>
        </div>

      </div>
      <div class="row">
        <div class="col-md-5 pr-md-1">
          <base-input label="Name"
                      v-model="testCollection.name">
          </base-input>
        </div>
      </div>
      <div class="row">
        <div class="col-md-6 pr-md-1">
          <base-input label="Description"
                      v-model="testCollection.description"
                      placeholder="">
          </base-input>
        </div>
      </div>
      <base-alert type="danger" v-if="showErrors">
        <span><b> Error - </b> {{ errorMessage }}</span>
      </base-alert>
      <base-button slot="footer" type="primary" @click="saveTestCollection" fill>Save</base-button>
    </card>

    <card>
      <template slot="header">
        <div class="row">
          <h6 class="title d-inline col-md-1">Tests</h6>
          <p class="card-category d-inline col-md-10">These are the tests</p>
          <base-dropdown menu-on-right=""
                         tag="div"
                         title-classes="btn btn-link btn-icon"
                         aria-label="Settings menu" class="col-md-1 align-right">
            <i slot="title" class="tim-icons icon-settings-gear-63"></i>
            <a class="dropdown-item" href="#/add-agent">Add new Test</a>
          </base-dropdown>
        </div>

      </template>
      <div class="table-responsive">
        <base-table :data="testCollection.tests"
                    thead-classes="text-primary">
          <template slot-scope="{row}">
            <td>
              <b>{{row.name}}</b>
            </td>
          </template>
        </base-table>
      </div>
    </card>
  </div>

</template>
<script>
import {data} from "@/shared/datastore.js";
import {BaseAlert, BaseTable} from '@/components'
import router from "@/router";

const uuid = require("uuid");

export default {
  name: "edit-testcollection",
  props: {
    testCollectionInput: {},
    testGroupIdInput: ''
  },
  components: {
    BaseAlert,
    BaseTable
  },
  data() {
    return {
      testCollection: {
        id: '',
        name: '',
        description: '',
        groupId: '',
        tests: []
      },
      type: ["", "info", "success", "warning", "danger"],
      showErrors: false,
      errorMessage: ''
    }
  },
  created() {
    console.log(this.testCollectionInput);
    console.log('TestGroupId input')
    console.log(this.testGroupIdInput);
    try {
      if (this.testCollectionInput.id !== "") {
        this.testCollection.id = this.testCollectionInput.id;
        this.testCollection.name = this.testCollectionInput.name;
        this.testCollection.description = this.testCollectionInput.description;
        this.testCollection.groupId = this.testCollectionInput.groupId;
        this.testCollection.tests = this.testCollectionInput.tests;
      }
    } catch (e) {
      // We are creating a new one so we are swallowing this
    }
    console.log(this.testCollection)
    if (this.testGroupIdInput){
      console.log('Executing GroupId update from input')
      this.testCollection.groupId = this.testGroupIdInput;
    }
  },
  methods: {
    async saveTestCollection() {
      this.showErrors = false;

      try {
        if (this.testCollection.name.length < 3) {
          this.showErrors = true;
          this.errorMessage = 'Name too short';
          return
        }
        if (this.testCollection.description.length < 3) {
          this.showErrors = true;
          this.errorMessage = 'Description too short';
          return
        }
      } catch (e) {
        this.showErrors = true;
        this.errorMessage = e;
        return
      }

      if (this.testCollection.id == '') {
        this.testCollection.id = uuid.v1();
      }

      let input = {
        id: this.testCollection.id,
        name: this.testCollection.name,
        description: this.testCollection.description,
        groupId: this.testCollection.groupId
      }
      try {
        await data.updateTestCollection(input);

        // router.push({name: 'testgroups'})
      } catch (e) {
        this.showErrors = true;
        this.errorMessage = e;
      }
    }
  }
}
</script>
<style>
</style>
