using System.Text.Json.Serialization;

var builder = WebApplication.CreateBuilder(args);

builder.Logging.AddConsole();
builder.Services.AddHealthChecks();

var app = builder.Build();

app.UseHealthChecks("/health");
app.MapGet("/api/v1/accounts/{id:int}", (int id) =>
{
    if (id % 2 != 0) return Results.NotFound();

    return Results.Ok(new Account(Guid.NewGuid(), Guid.NewGuid().ToString()));
});

app.Run();

internal record Account(Guid AccountId, string Owner);