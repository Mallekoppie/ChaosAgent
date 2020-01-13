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

export const data = {
  getAllTestGroups,
  updateTestCollection
};
