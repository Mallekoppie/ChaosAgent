const Agents = [
    {
        id: "someGuid",
        host: "localhost",
        port: 21000,
        enabled: true,
        status: "online"
    },
    {
        id: "someGuid1",
        host: "serverone",
        port: 21000,
        enabled: true,
        status: "online"
    },
    {
        id: "someGuid2",
        host: "serverTwo",
        port: 21000,
        enabled: true,
        status: "online"
    },
    {
        id: "someGuid3",
        host: "three",
        port: 21000,
        enabled: true,
        status: "online"
    },
    {
        id: "someGuid4",
        host: "four",
        port: 21000,
        enabled: true,
        status: "online"
    },
    {
        id: "someGuid5",
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
        description: "something that describes the group of test collections 1",
        testCollections: TestCollection,
        showDetail: false
    },
    {
        id: 2,
        name: "someTest 2",
        description: "something that describes the group of test collections 2",
        testCollections: TestCollection,
        showDetail: false
    },
    {
        id: 3,
        name: "someTest 3",
        description: "something that describes the group of test collections 3",
        testCollections: TestCollection,
        showDetail: false
    },
    {
        id: 4,
        name: "someTest 4",
        description: "something that describes the group of test collections 4",
        testCollections: TestCollection,
        showDetail: false
    },
    {
        id: 5,
        name: "someTest",
        description: "something that describes the group of test collections 5",
        testCollections: TestCollection,
        showDetail: false
    },
    {
        id: 6,
        name: "someTest",
        description: "something that describes the group of test collections 5",
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
            name: "test 2",
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
            name: "test 3",
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

const getTestCollections = function(){
    return TestCollection
}

export const dataStore = {
    getAgents,
    getTestGroups,
    getTestCollections,
    getTestDetail
};
