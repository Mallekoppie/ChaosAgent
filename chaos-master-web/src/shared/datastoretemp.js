const Agents = [
  {
    agentId: "someGuid",
    host: "localhost",
    port: 21000,
    enabled: true,
    status: "online"
  },
  {
    agentId: "someGuid1",
    host: "serverone",
    port: 21000,
    enabled: true,
    status: "online"
  },
  {
    agentId: "someGuid2",
    host: "serverTwo",
    port: 21000,
    enabled: true,
    status: "online"
  },
  {
    agentId: "someGuid3",
    host: "three",
    port: 21000,
    enabled: true,
    status: "online"
  },
  {
    agentId: "someGuid4",
    host: "four",
    port: 21000,
    enabled: true,
    status: "online"
  },
  {
    agentId: "someGuid5",
    host: "five",
    port: 21000,
    enabled: true,
    status: "online"
  }
];

const getAgents = function() {
  return Agents;
};

const TestCollection = [
  {
    id: 1,
    name: "some name",
    description: "lekker description",
    Tests: []
  },
  {
    id: 2,
    name: "some name",
    description: "lekker description",
    Tests: []
  },
  {
    id: 3,
    name: "some name",
    description: "lekker description",
    Tests: []
  },
  {
    id: 4,
    name: "some name",
    description: "lekker description",
    Tests: []
  }
];

const TestGroups = [
  {
    id: 1,
    name: "someTest",
    description: "something that describes the group of test collections",
    testCollections: TestCollection,
    showDetail: false
  },
  {
    id: 2,
    name: "someTest",
    description: "something that describes the group of test collections",
    testCollections: TestCollection,
    showDetail: false
  },
  {
    id: 3,
    name: "someTest",
    description: "something that describes the group of test collections",
    testCollections: TestCollection,
    showDetail: false
  },
  {
    id: 4,
    name: "someTest",
    description: "something that describes the group of test collections",
    testCollections: TestCollection,
    showDetail: false
  },
  {
    id: 5,
    name: "someTest",
    description: "something that describes the group of test collections",
    testCollections: TestCollection,
    showDetail: false
  },
  {
    id: 6,
    name: "someTest",
    description: "something that describes the group of test collections",
    testCollections: TestCollection,
    showDetail: false
  }
];

const getTestGroups = function() {
  return TestGroups;
};

export const dataStore = {
  getAgents,
  getTestGroups
};
