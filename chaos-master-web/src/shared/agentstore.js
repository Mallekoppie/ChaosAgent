import * as axios from "axios";

import { API } from "./config";

const getAllAgents = async function() {
  console.log(`API from environment: ${API}`);

  const response = await axios.get(`${API}/agents`);

  if (response.status !== 200) {
    throw Error("Error retrieving all test groups");
  }

  let agents = response.data;

  return agents;
};

const deleteAgent = async function(agent) {
  const response = await axios.delete(
    `${API}/agents`,
    { data: agent },
    {
      headers: { "Content-Type": "application/json" }
    }
  );

  if (response.status !== 200) {
    throw Error("Error retrieving all test groups");
  }
};

const updateAgent = async function(agent) {
  console.log("Executing agent update");
  const response = await axios.put(`${API}/agents`, agent, {
    headers: { "Content-Type": "application/json" }
  });

  if (response.status !== 200) {
    throw Error(
      `Updating of Test Collection failed. ResponseCode: ${response.statusText}`
    );
  }
};

export const data = {
  getAllAgents,
  deleteAgent,
  updateAgent
};
