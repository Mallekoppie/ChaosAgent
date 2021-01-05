<template>
  <card>
    <h5 slot="header" class="title">Edit Agent</h5>
    <div class="row">
      <div class="col-md-5 pr-md-1">
        <base-input label="ID (generated)"
                    placeholder="id"
                    v-model="agent.id"
                    disabled>
        </base-input>
      </div>
      <div class="col-md-3 px-md-1">
        <base-input label="Host"
                    placeholder="host"
                    v-model="agent.host">
        </base-input>
      </div>
    </div>
    <div class="row">
      <div class="col-md-6 pr-md-1">
        <base-input label="Port"
                    v-model="agent.port"
                    placeholder="0">
        </base-input>
      </div>
      <div class="col-md-6 pl-md-1">
        <base-input label="Metrics Port"
                    v-model="agent.metricsPort"
                    placeholder="0">
        </base-input>
      </div>
    </div>
    <div class="row">
      <div class="col-md-4 pr-md-1">
        <base-input label="Status"
                    v-model="agent.status"
                    placeholder="">
        </base-input>
      </div>
    </div>
    <div class="col-md-4">
      <base-checkbox v-model="agent.enabled">
        Enabled
      </base-checkbox>
    </div>
    <base-alert type="danger" v-if="showErrors">
      <span><b> Error - </b> {{ errorMessage }}</span>
    </base-alert>
    <base-button slot="footer" type="primary" @click="saveAgent" fill>Save</base-button>
  </card>

</template>
<script>
import {data} from "@/shared/agentstore.js";
import {BaseAlert} from '@/components'
import router from "@/router";

const uuid = require("uuid");

export default {
  name: "Add-Agent",
  props: {
    agentInput: {}
  },
  components: {
    BaseAlert
  },
  data() {
    return {
      agent: {
        id: '',
        host: 0,
        port: 0,
        metricsPort: '',
        enabled: false,
        status: ''
      },
      type: ["", "info", "success", "warning", "danger"],
      showErrors: false,
      errorMessage: ''
    }
  },
  created() {
    if (this.agentInput.id !== "") {
      this.agent.id = this.agentInput.id;
      this.agent.host = this.agentInput.host;
      this.agent.port = this.agentInput.port;
      this.agent.metricsPort = this.agentInput.metricsPort;
      this.agent.enabled = this.agentInput.enabled;
      this.agent.status = this.agentInput.status;
    }
  },
  methods: {
    async saveAgent() {
      this.showErrors = false;

      try {
        if (isNaN(this.agent.port) || parseInt(this.agent.port) < 1024) {
          this.showErrors = true;
          this.errorMessage = 'Port must be greater than 1024';
          return
        }
        if (isNaN(this.agent.metricsPort) || parseInt(this.agent.metricsPort) < 1024) {
          this.showErrors = true;
          this.errorMessage = 'Metrics Port must be greater than 1024';
          return
        }
        if (this.agent.host.length < 3) {
          this.showErrors = true;
          this.errorMessage = 'Host name too short';
          return
        }
      } catch (e) {
        this.showErrors = true;
        this.errorMessage = e;
        return
      }

      if (this.agent.id == '') {
        this.agent.id = uuid.v1();
      }

      let input = {
        id: this.agent.id,
        host: this.agent.host,
        port: parseInt(this.agent.port),
        metricsPort: parseInt(this.agent.metricsPort),
        enabled: this.agent.enabled,
        status: this.agent.status
      }
      try {
        await data.updateAgent(input);

        router.push({name:'agents'})
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
