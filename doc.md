# toy robot in go

## First design

app sends input to parser which gives back commands; app then executes those commands on its own model

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
 
app sends input to parser which gives back commands. app has model stored in state. App passes state and commands to handler which executes the command on the state, updating the state's model.

```graphviz
digraph {
    { move place } -> { robot table }
    { turn report } -> { robot }
    { app } -> { state parser command handler }
    { parser } -> { command }
    { handler } -> { move place turn report }
    handler -> command
    handler -> state
    { state } -> { robot table }
}
```

## Third design: fewest abstractions

```graphviz
digraph {
    main -> {Scan Handle}
    Handle -> {Robot Table Sscanf Facing ParseFacing Place Turn Move Printf}
    ParseFacing -> Facing
    { Move Place } -> { Robot Table ValidPosition }
    ValidPosition -> Table
    Turn -> { Robot Facing }
    main -> { Robot Table }
    Robot -> Facing
}
```

Integration tests would be easy with this design, just feed input lines to Handle and then check model.

## Fourth design: command pattern

```graphviz
digraph {
    main -> { Scan ProcessLine }
    ProcessLine -> { Parse Command Model } 
    Model -> { Robot Table }
    Parse -> Command
    Command -> { Robot Table Sscanf Facing ParseFacing Place Turn Move Printf }
    ParseFacing -> Facing
    { Move Place } -> { Robot Table ValidPosition }
    Turn -> { Robot Facing }
    main -> { Robot Table }
}
```

How to decide which command to call? Pull out first word of command and select? Give line to each command until one can parse it? Yeah, that's better (but possibly less performant).

Update go to 1.12.5 using `snap install go --classic`.

Try again with packages.

Now that I've finished coding, redraw architectural diagram.

```graphviz
digraph {
    app -> { model commands command_Parse command_Execute }
    model -> { model_robot model_table }
    model_robot -> { model_facing }
    ValidPosition -> { model_table }
    commands -> { command_move command_place command_report command_turn }
    {command_move command_place } -> { command_Parse command_Execute ValidPosition model_robot model_table model_facing }
    command_report -> { model_robot command_Parse command_Execute }
    command_turn -> { model_robot model_facing command_Parse command_Execute }
    
}
```

250 lines of code, 250 lines of tests.

Diagram looks pretty overcomplicated :( I can't see a way of simplifying the diagram (by changing the code architecture).

Maybe if I take out the Parse and Execute?

```graphviz
digraph {
    app -> { model commands }
    model -> { model_robot model_table }
    model_robot -> { model_facing }
    ValidPosition -> { model_table }
    commands -> { command_move command_place command_report command_turn }
    {command_move command_place } -> { ValidPosition model_robot model_table model_facing }
    command_report -> { model_robot }
    command_turn -> { model_robot model_facing }
    
}
```

That looks clearer... I can't see a better way. Needs a few integration tests, maybe:

- PLACE invalid, REPORT
- PLACE, REPORT, TURN, MOVE, REPORT
- several invalid commands then PLACE then some more invalid commands then REPORT

