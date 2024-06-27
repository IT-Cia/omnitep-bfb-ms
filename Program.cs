using Microsoft.AspNetCore.Server.Kestrel.Https;
using omnitep_bfb_ms.Services;
using System.Configuration;
using System.Net;

var builder = WebApplication.CreateBuilder(args);

builder.WebHost.UseKestrel(options =>
{
    options.ConfigureHttpsDefaults(o =>
    {
        //o.ClientCertificateValidation += ValidateClientCertificate;
        o.AllowAnyClientCertificate();
        o.ClientCertificateMode = ClientCertificateMode.RequireCertificate;
        o.CheckCertificateRevocation = false;
    });
    options.Listen(IPAddress.Loopback, builder.Configuration.GetValue<int>("PortNumber"));
    options.Listen(IPAddress.Loopback, builder.Configuration.GetValue<int>("PortNumberSSL"));
});

var httpHandler = new HttpClientHandler
{
    ServerCertificateCustomValidationCallback = HttpClientHandler.DangerousAcceptAnyServerCertificateValidator
};

builder.Services.AddGrpcClient<Notifications.notifications.notificationsClient>(o =>
{
    o.Address = new Uri(builder.Configuration.GetValue<string>("GrpcUrl"));
}).ConfigureChannel(o =>
{
    o.HttpHandler = httpHandler;
});

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddSingleton<GrpcService>();

var app = builder.Build();

// Configure the HTTP request pipeline.
//if (app.Environment.IsDevelopment())
//{
    app.UseSwagger();
    app.UseSwaggerUI();
//}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();