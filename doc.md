# toy robot in go

## First design

app sends input to parser which gives back commands; app then executes those commands on its own state

```graphviz
digraph {
    { move place } -> { robot table }
    { turn report } -> { robot }
    { app } -> { parser command robot table }
    { parser } -> { command }
    { command } -> { move place turn report }
}
```

## Second design
 
app sends input to parser which gives back commands. app has state stored in model. App passes model and commands to handler which executes the command on the model, updating the model's state.

```graphviz
digraph {
    { move place } -> { robot table }
    { turn report } -> { robot }
    { app } -> { model parser command handler }
    { parser } -> { command }
    { handler } -> { move place turn report }
    handler -> command
    handler -> model
    { model } -> { robot table }
}
```