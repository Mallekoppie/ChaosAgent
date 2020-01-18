import * as axios from "axios";

import { API } from "./config";

const getAllTestGroups = async function() {
  console.log(`API from environment: ${API}`);

  const response = await axios.get(`${API}/testgroups`);

  if (response.status !== 200) {
    throw Error("Error retrieving all test groups");
  }

  let tests = response.data;

  tests.forEach(function(item) {
    item.showDetail = false;
  });

  return tests;
};

const getTestGroup = async function(id) {
  console.log(`API from environment: ${API}`);

  const response = await axios.get(`${API}/testgroups/` + id);

  if (response.status !== 200) {
    throw Error("Error retrieving all test groups");
  }

  let group = response.data;

  group.showDetail = false;

  return group;
};

const addTestGroup = async function(testGroup) {
  console.log("Executing test collection update");
  const response = await axios.post(`${API}/testgroups`, testGroup, {
    headers: { "Content-Type": "application/json" }
  });

  if (response.status !== 201) {
    throw Error(
      `Updating of Test Collection failed. ResponseCode: ${response.statusText}`
    );
  }
};

const deleteTestGroup = async function(id) {
  console.log("Executing delete test group");
  const response = await axios.delete(`${API}/testgroups/` + id);

  if (response.status !== 204) {
    throw Error(
      `Updating of Test Collection failed. ResponseCode: ${response.statusText}`
    );
  }
};

const updateTestCollection = async function(testCollection) {
  console.log("Executing test collection update");
  const response = await axios.put(`${API}/testcollections`, testCollection, {
    headers: { "Content-Type": "application/json" }
  });

  if (response.status !== 201) {
    throw Error(
      `Updating of Test Collection failed. ResponseCode: ${response.statusText}`
    );
  }
};

const deleteTestCollection = async function(id) {
  console.log("Executing delete test collection for id: " + id);
  const response = await axios.delete(`${API}/testcollections/` + id);

  if (response.status !== 204) {
    throw Error(
      `Updating of Test Collection failed. ResponseCode: ${response.statusText}`
    );
  }
};

export const data = {
  getAllTestGroups,
  getTestGroup,
  updateTestCollection,
  deleteTestCollection,
  addTestGroup,
  deleteTestGroup
};
