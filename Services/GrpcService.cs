using Grpc.Net.Client;
using System.Net;

namespace omnitep_bfb_ms.Services
{
    public class GrpcService
    {
        //private readonly Notifications.notifications.notificationsClient client;

        public GrpcService(IConfiguration configuration)
        {
           // HttpClient.DefaultProxy = new WebProxy();

            //var channel = GrpcChannel.ForAddress(configuration.GetValue<string>("GrpcUrl"));
            //client = new Notifications.notifications.notificationsClient(channel);
        }

        //public Notifications.notifications.notificationsClient GetClient()
        //{
            //return client;
        //}
    }
}
