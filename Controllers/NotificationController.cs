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

        public NotificationController(GrpcService grpc, IConfiguration configuration)
        {
            this.grpc = grpc;
            this.configuration = configuration;
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
            
            var response = await grpc.GetClient().SendAsync(request);
            return new OkObjectResult(response);
            //}
            //catch (RpcException ex)Ö
            //{
            //    return new BadRequestObjectResult(ex.Status.StatusCode + ": " + ex.Status.Detail);
            //}
            
        }
    }
}