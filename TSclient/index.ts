import { FetchRequestAdapter } from "@microsoft/kiota-http-fetchlibrary";
import { createApiClient } from "./apiClient";
import { AnonymousAuthenticationProvider, UntypedNode, isUntypedArray, isUntypedObject } from "@microsoft/kiota-abstractions";


// API requires no authentication, so use the anonymous
// authentication provider
const authProvider = new AnonymousAuthenticationProvider();
// Create request adapter using the fetch-based implementation
const adapter = new FetchRequestAdapter(authProvider);
// Create the API client
const client = createApiClient(adapter);

async function main(): Promise<void> {
    try {
        // GET /metrics.json
        const allPosts = await client.metricsJson.get();

        const dataSets = allPosts?.datasets;
        parseUntypedJson(dataSets);

    } catch (err) {
        console.log(err);
    }
}

function parseUntypedJson(dataSets: UntypedNode | undefined, indent: string = "") {
    if(dataSets === undefined) {
        return;
    }

    if(isUntypedArray(dataSets)) {
        console.log(indent+ "Found array type: ");
        for(const dataSet of dataSets.getValue()) {
            parseUntypedJson(dataSet, indent + "  ");
        }
    }
    else if (isUntypedObject(dataSets)) {
        const properties = dataSets.getValue();
        console.log(indent + "Found object type: ");
        for (const [key,value] of Object.entries(properties)) {
            console.log(indent + "Found property with name: " + key);
            parseUntypedJson(value, indent + "  ");
        }
    }
    else {
        console.log(indent + "Found scalar value: " + dataSets.getValue());
    }
}

// run it!
main()