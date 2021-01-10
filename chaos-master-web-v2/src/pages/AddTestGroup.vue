<template>
  <card>
    <h5 slot="header" class="title">Edit Agent</h5>
    <div class="row">
      <div class="col-md-5 pr-md-1">
        <base-input label="ID (generated)"
                    placeholder="id"
                    v-model="testgroup.id"
                    disabled>
        </base-input>
      </div>
      <div class="col-md-3 px-md-1">
        <base-input label="Name"
                    v-model="testgroup.name">
        </base-input>
      </div>
    </div>
    <div class="row">
      <div class="col-md-6 pr-md-1">
        <base-input label="Description"
                    v-model="testgroup.description"
                    placeholder="">
        </base-input>
      </div>
    </div>
    <base-alert type="danger" v-if="showErrors">
      <span><b> Error - </b> {{ errorMessage }}</span>
    </base-alert>
    <base-button slot="footer" type="primary" @click="saveTestGroup" fill>Save</base-button>
  </card>

</template>
<script>
import {data} from "@/shared/datastore.js";
import {BaseAlert} from '@/components'
import router from "@/router";

const uuid = require("uuid");

export default {
  name: "Add-TestGroup",
  props: {
    testgroupInput: {}
  },
  components: {
    BaseAlert
  },
  data() {
    return {
      testgroup: {
        id: '',
        name: '',
        description: ''
      },
      type: ["", "info", "success", "warning", "danger"],
      showErrors: false,
      errorMessage: ''
    }
  },
  created() {
    if (this.testgroupInput.id !== "") {
      this.testgroup.id = this.testgroupInput.id;
      this.testgroup.name = this.testgroupInput.name;
      this.testgroup.description = this.testgroupInput.description;
    }
  },
  methods: {
    async saveTestGroup() {
      this.showErrors = false;

      try {
        if (this.testgroup.name.length < 3) {
          this.showErrors = true;
          this.errorMessage = 'Name too short';
          return
        }
        if (this.testgroup.description.length < 3) {
          this.showErrors = true;
          this.errorMessage = 'Description too short';
          return
        }
      } catch (e) {
        this.showErrors = true;
        this.errorMessage = e;
        return
      }

      if (this.testgroup.id == '') {
        this.testgroup.id = uuid.v1();
      }

      let input = {
        id: this.testgroup.id,
        name: this.testgroup.name,
        description: this.testgroup.description
      }
      try {
        await data.updateTestGroup(input);

        router.push({name:'testgroups'})
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
