import { FetchRequestAdapter } from "@microsoft/kiota-http-fetchlibrary";
import { createApiClient } from "./apiClient";
import { AnonymousAuthenticationProvider, UntypedNode } from "@microsoft/kiota-abstractions";


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

function parseUntypedJson(dataSets: UntypedNode) {
    console.log("dataSets");
}
