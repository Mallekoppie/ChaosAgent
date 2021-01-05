<template>
  <card>
    <h5 slot="header" class="title">Edit Agent</h5>
    <div class="row">
      <div class="col-md-5 pr-md-1">
        <base-input label="ID (disabled)"
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
        <base-input label="Enabled"
                    v-model="agent.enabled"
                    placeholder="false">
        </base-input>
      </div>
      <div class="col-md-4 px-md-1">
        <base-input label="Status"
                    v-model="agent.status"
                    placeholder="">
        </base-input>
      </div>
    </div>
    <base-button slot="footer" type="primary" @click="saveAgent" fill>Save</base-button>
  </card>

</template>
<script>
import { data } from "@/shared/agentstore.js";
const uuid = require("uuid");

  export default {
    name: "Add-Agent",
    props: {
      agentInput:{}
    },
    components: {
    },
    data() {
      return {
        agent: {
          id: '',
          host: 0,
          port: 0,
          metricsPort: '',
          enabled: '',
          status: ''
        },
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
        if (this.agent.id == '') {
          this.agent.id = uuid.v1();
        }

        let input = {
          id: this.agent.id,
          host: this.agent.host,
          port: parseInt(this.agent.port),
          metricsPort: parseInt(this.agent.metricsPort),
          enabled: (this.agent.enabled == 'true'),
          status: this.agent.status
        }
        await data.updateAgent(input);
      }
    }
  }
</script>
<style>
</style>
