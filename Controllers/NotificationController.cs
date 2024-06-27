using Grpc.Core;
using Grpc.Net.Client;
using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Mvc;
using omnitep_bfb_ms.Services;
using Swashbuckle.AspNetCore.Annotations;
using System.ComponentModel.DataAnnotations;
using System.Configuration;

namespace omnitep_bfb_ms.Controllers
{
    [ApiController]
    [Microsoft.AspNetCore.Mvc.Route("api/[controller]")]
    [SwaggerTag("SwaggerTag on class")]
    public class NotificationController : ControllerBase
    {
        private GrpcService grpc;
        private IConfiguration configuration;
        private Notifications.notifications.notificationsClient client;

        public NotificationController(GrpcService grpc, IConfiguration configuration, Notifications.notifications.notificationsClient client)
        {
            this.grpc = grpc;
            this.configuration = configuration;
            this.client = client;
        }

        [HttpPost]
        [Consumes("application/json")]
        [Produces("application/json")]
        public async Task<ActionResult<Notifications.SendResponse>> Post([Required] Notifications.SendRequest request)
        {
            if (string.IsNullOrWhiteSpace(request.From))
                return new BadRequestObjectResult("Parameter 'from' is required");
            if (string.IsNullOrWhiteSpace(request.To))
                return new BadRequestObjectResult("Parameter 'to' is required");
            if (string.IsNullOrWhiteSpace(request.Channel))
                return new BadRequestObjectResult("Parameter 'channel' is required");
            if (string.IsNullOrWhiteSpace(request.Message))
                return new BadRequestObjectResult("Parameter 'message' is required");

            //try
            //{
            
            var channel = GrpcChannel.ForAddress(configuration.GetValue<string>("GrpcUrl"));
            var client = new Notifications.notifications.notificationsClient(channel);
            var response = await client.SendAsync(request);
            return new OkObjectResult(response);
            //}
            //catch (RpcException ex)�
            //{
            //    return new BadRequestObjectResult(ex.Status.StatusCode + ": " + ex.Status.Detail);
            //}
            
        }
    }
}