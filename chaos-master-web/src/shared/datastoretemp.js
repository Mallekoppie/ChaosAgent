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

const TestDetail = {
  id: 1,
  name: "some name",
  description: "lekker description",
  tests: [
    {
      id: 1,
      name: "test 1",
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
    },
    {
      id: 2,
      name: "test 1",
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
    },
    {
      id: 3,
      name: "test 1",
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
    }
  ]
};

const getTestDetail = function() {
  return TestDetail;
};

export const dataStore = {
  getAgents,
  getTestGroups,
  getTestDetail
};
