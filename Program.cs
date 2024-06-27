using Microsoft.AspNetCore.Server.Kestrel.Https;
using omnitep_bfb_ms.Services;
using System.Configuration;
using System.Net;

var builder = WebApplication.CreateBuilder(args);

builder.WebHost.UseKestrel(options =>
{
    options.ConfigureHttpsDefaults(o =>
    {
        o.AllowAnyClientCertificate();
        o.ClientCertificateMode = ClientCertificateMode.RequireCertificate;
        o.CheckCertificateRevocation = false;
    });
});

/*var httpHandler = new HttpClientHandler
{
    ServerCertificateCustomValidationCallback = HttpClientHandler.DangerousAcceptAnyServerCertificateValidator
};

builder.Services.AddGrpcClient<Notifications.notifications.notificationsClient>(o =>
{
    o.Address = new Uri(builder.Configuration.GetValue<string>("GrpcUrl"));
}).ConfigureChannel(o =>
{
    o.HttpHandler = httpHandler;
});*/

builder.Services.AddControllersWithViews();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddSingleton<GrpcService>();

builder.WebHost.UseUrls(Environment.GetEnvironmentVariable("APP_URL") ?? "http://localhost:5000");

var app = builder.Build();

// Configure the HTTP request pipeline.
//if (app.Environment.IsDevelopment())
//{
    app.UseSwagger();
    app.UseSwaggerUI();
//}

//app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
