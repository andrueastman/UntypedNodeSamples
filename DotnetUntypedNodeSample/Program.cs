using Graphdotnetv4;
using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Abstractions.Serialization;
using Microsoft.Kiota.Http.HttpClientLibrary;
using System.Xml.Linq;

namespace DotnetUntypedNodeSample
{
    internal class Program
    {
        static async Task Main(string[] args)
        {
            var requestAdapter = new HttpClientRequestAdapter(new AnonymousAuthenticationProvider());
            var apiGuruClient = new ApiClient(requestAdapter);

            var metrics = await apiGuruClient.MetricsJson.GetAsync();
            // print out datasets property which is untyped json
            var dataSets = metrics?.Datasets;
            ParseUnknownObject(dataSets);
        }
        private static void ParseUnknownObject(UntypedNode untypedNode, string indent = "")
        {
            switch (untypedNode)
            {
                case UntypedObject untypedObject:
                    Console.WriteLine(indent + "Found object value: ");
                    var properties = untypedObject.GetValue();
                    foreach (var (name, node) in properties)
                    {
                        Console.WriteLine(indent + "Property Name: " + name);
                        ParseUnknownObject(node, indent + "  ");
                    }
                    break;
                case UntypedArray untypedArray:
                    Console.WriteLine(indent + "Found array value: ");
                    foreach (var item in untypedArray.GetValue())
                    {
                        Console.WriteLine(indent + "New Item: ");
                        ParseUnknownObject(item, indent + "  ");
                    }
                    break;
                default:
                    Console.WriteLine(indent + "Found scalar value : " + KiotaJsonSerializer.SerializeAsString(untypedNode));
                    break;

            }

        }
    }
}
