
<div align="center">
  <img width="50%" alt="roman men playing tic-tac-toe" src="../assets/TTTLogo.png" />
</div>

# 3T Project Documentation




## Table of content
1. [Game State](#game-state)
   1. [State Representation](#state-representation)
      1. [Approaches](#approaches)
      2. [Final Representation](#bit-layout-left--right)
2. [Game Graph/Map](#game-graph)
   1. [Graph/Map Representation](#game-graph-representation)
   2. [Next States Generation](#next-state-generation)
      1. [Symmetry Reduction](#symmetry-reduction)
            - [Rotation Symmetry](#1-rotational-symmetry)
            - [Reflection Symmetry](#2-reflection-symmetry)
            - [Cannonical State](#canonical-state)
   3. [Setting Scores](#setting-scores)
   4. [MaxiMax: Self-proclaimed MiniMax varient](#maximax-algorithm)
3. [Game Engine](#game-engine)



## Game State
A **Game State** represents a complete snapshot of the Tic-Tac-Toe board at a specific moment in time.

Each state is:
- Immutable
- Self-contained
- Uniquely identifiable

A state fully defines:
- Board configuration

This design enables deterministic graph traversal, serialization.


### State Representation
For the state representation, we require to only represent the occupied position of X and O in the 9 cells 
of Tic-Tac-Toe game board.

At any given point in time, a cell can be either:
- occupied by x
- occupied by O
- vacant


### Approaches
1. [Array of length 9](#approach-1-array-of-length-9)
2. [Integer with 9 digits](#approach-2-9-digits-integer)
3. [32 bit unsigned integer](#approach-3-32-bit-unsigned-integer)



### Approach 1: Array of length 9
If we represent a cell:
- occupied by x with 1
- occupied by O with 2
- vacant with 0

and use an array of length 9, it will occupy size $= 9*sizeOf(each\_element) = 9*8 = 72\ bytes = 576\ bits$
>\* since the smallest int type in golang requires size 8 bytes

### Approach 2: 9 digits Integer
If we represent a cell:
- occupied by x with 1
- occupied by O with 2
- vacant with 0

and use an integer of 9 digits, it will occupy size $= log_2(10^{9+1})\ bits = 33.22\ bits \approx 34\ bits$

>\* In golang it will need 64 bits at least


### Approach 3: 32-bit unsigned integer
I tried a lot of ways to represent the game state that would require minimum resource, and came up 
with a representation that requires only 18 bits:
- 9 bits to represent cells occupied by X
- 9 bits to represent cells occupied by O

This requires 18 bits and in go will need 32 bits with uint32 or int32 type

Noticing that we still have 14 bits free in out 32 bits variable (if used uint32), to reduce the computation
for:
- vacancy cells (we can store 9 bits for in the number itself) 
- Whether the game has ended or not? (we can store 1 bit for this) 
- - If ended then is it a tie or not? (we can store 1 bit for this)
- - - If not tie then who won? (we can store 1 bit for this)
- - - If not tie then who lost? (we can store 1 bit for this)
- If not ended then who will play next turn? (we can store 1 bit for this)

So, we utilize the 32 bit resource completely.

### Final State representation
### Bit Layout (Left → Right)

### Field Description

<div align="center">

| Field | Bits | Meaning                                           |
|------|------|---------------------------------------------------|
| `C` | 1 | Continuum flag (not a terminal state)             |
| `T` | 1 | Tie state  (is the state a tie terminal state)    |
| `W` | 1 | Win state  (who won this state : (0 = X, 1 = O))  |
| `L` | 1 | Loss state (who lost this state : (0 = X, 1 = O)) |
| `Turn` | 1 | Next player (0 = X, 1 = O)                        |
| `X` | 9 | Bitmask of X positions (0 = absent, 1 = present)  |
| `O` | 9 | Bitmask of O positions (0 = absent, 1 = present)                          |
| `Vacancy` | 9 | Empty board positions (0 = absent, 1 = present) |

</div>

- Each board cell is represented by a single bit.
- For any cell, exactly one of `X`, `O`, or `Vacancy` is set.
- No redundant information is stored.

>The entire game state is encoded into a single **32-bit unsigned integer (`uint32`)** called `State`.

```go
type State uint32
```
>declared [here](../src/model/structure.go#L7)
## Game Graph

The game graph would be a uni-directional graph connecting the nodes (States) to its children nodes (next States)

### Game graph representation
The game graph is a key value pair map with key of type State (uint32) 
and the value of type struct StateProps (short for State Properties)

State Properties:
* NextStates: A slice []State containing all the next states 
* Score of type float32 to store score of a state as a move
* WinDepth: the maximum number of steps in which the State as a move guarantees a win if played strategically, value 10 means (cannot guarantee a win from this state)
* LoseDepth: the minimum number of steps in which the State as a move guarantees a lose if the opponent play optimal, value 10 means (cannot guarantee a lose from this state)

```go
type StateProps struct {
	Score      float32
	WinDepth   int8
	LoseDepth  int8
	NextStates []State
}

type GMap map[State]StateProps

```
>declared [here](../src/model/map.go#L11-L18)

### Next State Generation

>The function `NextMoves` which can be found [here](../src/utils/nextstates.go#L11-L40) </br>
> uses bit-Manipulation to generate a slice containing next states, each of type State</br>

### Symmetry reduction

However, we need to realize that there exists symmetry in tic-tac-toe board.


### 1. Rotational Symmetry

<table align="center">
  <tr>
    <td align="center">0°</td>
    <td align="center">90°</td>
    <td align="center">180°</td>
    <td align="center">270°</td>
  </tr>
  <tr>
    <td>
      <table>
        <tr><td>X</td><td>O</td><td>ㅤ</td></tr>
        <tr><td>ㅤ</td><td>ㅤ</td><td>ㅤ</td></tr>
        <tr><td>ㅤ</td><td>X</td><td>ㅤ</td></tr>
      </table>
    </td>
    <td>
      <table>
        <tr><td>ㅤ</td><td>ㅤ</td><td>X</td></tr>
        <tr><td>X</td><td>ㅤ</td><td>O</td></tr>
        <tr><td>ㅤ</td><td>ㅤ</td><td>ㅤ</td></tr>
      </table>
    </td>
    <td>
      <table>
        <tr><td>ㅤ</td><td>X</td><td>ㅤ</td></tr>
        <tr><td>ㅤ</td><td>ㅤ</td><td>ㅤ</td></tr>
        <tr><td>ㅤ</td><td>O</td><td>X</td></tr>
      </table>
    </td>
    <td>
      <table>
        <tr><td>ㅤ</td><td>ㅤ</td><td>ㅤ</td></tr>
        <tr><td>O</td><td>ㅤ</td><td>X</td></tr>
        <tr><td>X</td><td>ㅤ</td><td>ㅤ</td></tr>
      </table>
    </td>
  </tr>
</table>


### 2. Reflection Symmetry

<table align="center">
  <tr>
    <td align="center">Normal</td>
    <td align="center">Horizontal Flip</td>
    <td align="center">Vertical Flip</td>
    <td align="center">Diagonal Flip</td>
    <td align="center">Anti-Diagonal Flip</td>
  </tr>
  <tr>
    <td>
      <table>
        <tr><td>X</td><td>O</td><td>ㅤ</td></tr>
        <tr><td>ㅤ</td><td>ㅤ</td><td>ㅤ</td></tr>
        <tr><td>ㅤ</td><td>X</td><td>ㅤ</td></tr>
      </table>
    </td>
    <td>
      <table>
        <tr><td>ㅤ</td><td>O</td><td>X</td></tr>
        <tr><td>ㅤ</td><td>ㅤ</td><td>ㅤ</td></tr>
        <tr><td>ㅤ</td><td>X</td><td>ㅤ</td></tr>
      </table>
    </td>
    <td>
      <table>
        <tr><td>ㅤ</td><td>X</td><td>ㅤ</td></tr>
        <tr><td>ㅤ</td><td>ㅤ</td><td>ㅤ</td></tr>
        <tr><td>X</td><td>O</td><td>ㅤ</td></tr>
      </table>
    </td>
    <td>
      <table>
        <tr><td>X</td><td>ㅤ</td><td>ㅤ</td></tr>
        <tr><td>O</td><td>ㅤ</td><td>X</td></tr>
        <tr><td>ㅤ</td><td>ㅤ</td><td>ㅤ</td></tr>
      </table>
    </td>
    <td>
      <table>
        <tr><td>ㅤ</td><td>ㅤ</td><td>X</td></tr>
        <tr><td>ㅤ</td><td>ㅤ</td><td>ㅤ</td></tr>
        <tr><td>X</td><td>O</td><td>ㅤ</td></tr>
      </table>
    </td>
  </tr>
</table>

$\therefore$ there are total 8 symmetries possible:

1. Normal 
2. 90° Rotation [find in code](../src/Constants/symmetryTables.go#L3-L36)
3. 180° Rotation [find in code](../src/Constants/symmetryTables.go#L38-L71)
4. 270° Rotation [find in code](../src/Constants/symmetryTables.go#L73-L106)
5. Horizontal Flip [find in code](../src/Constants/symmetryTables.go#L108-L141)
6. Vertical Flip [find in code](../src/Constants/symmetryTables.go#L143-L176)
7. Diagonal Flip [find in code](../src/Constants/symmetryTables.go#L178-L211)
8. Anti-Diagonal Flip [find in code](../src/Constants/symmetryTables.go#L213-L246)


### Canonical State
This means in order to remove symmetry redundancy, we need to collapse all the symmetric 
forms to one single form. But how to decide which one to select for all the 8 forms
I chose the one which translates smallest number since all states are binary number after all.

Using this I generated the game graph which contains 765 states.</br>
Out of these 765 states:</br>
X wins in 91 states</br>
O wins in 44 states

### Setting Scores
#### Initial Settings
At the time of game map generation, I set:</br>
score = 0.0</br>
WinDepth = 10</br>
LoseDepth = 10</br>
#### Setting terminal values
For the terminal winning states (i.e. all states that are terminal and non-tie states), set the values:</br>
score -> $\infty$</br>
WinDepth -> 0</br>

#### MaxiMax algorithm
Applied the `MaxiMax` (Self proclaimed variant of MiniMax algorithm, to be applied on full game graphs and 
not just limited to trees) [find in code](../src/mapBuilder/setValues.go#L37-L97) 

```
workedHistory = []
winStates = getTerminalWinningStates(gameMap)
loseStates = []

wins():
    if len(winStates) == 0:
        return
    
    loseArray = []
    
    for node in winStates:
        node.score = infinity
        if len(children(node)) != 0:
            node.Windepth = max(children(node).LoseDepth)
        else:
            node.WinDepth = 0
        
        for parent in parents(node):
            if parent not in workedHistory:
                parent.score = -infinity
                parent.Losedepth = min(children(parent).WinDepth)+1
            workedHistory.insert(parent)
    
    loses()
    return
    
loses():
    if len(loseStates) == 0:
        return
    winStates = []
    
    for node in loseStates:
        for parent in parents(node):
            if parent not in workedhistory:
                if all(children(parent).score) == -infinity:
                    winStates.insert(parent)
    
    wins()
    return   
```

>*Note:*</br>
> I also added a custom function to score the states which cannot guarantee a win
so there score would have remained 0.0 I gave them score with the average score (or simply average positive score) of their children.

## Game Engine

Once the game map values have been set. Making a game engine is no big deal.
Just make sure load the map file before the `TTTNGinPlay` [find in code](../src/TTTEngine/3TEngine.go#L143-L145)
function is called to avoid reloading everytime (huge wastage of memory & time) the function is called.

1. For a given move, convert the move into its canonical form.
2. Get the moves that can be played by using gameMap[state].NextStates
3. Extract those moves which have the highest score among them [find in code](../src/TTTEngine/3TEngine.go#L23-L31)
4. If score is infinity, select the move which is minimum WinDepth (defeat opponent in the least possible steps) [find in code](../src/TTTEngine/3TEngine.go#L46-L59)
5. If score is -infinity, select the move which is maximum loseDepth (increase the gameplay so, as to have more chances of opponent making a mistake) [find in code](../src/TTTEngine/3TEngine.go#L33-L46)
6. If the score is other floating number, select a random one among them. [find in code](../src/TTTEngine/3TEngine.go#L59-L63)
7. Now we got the canonical form of the best move to play, we still need to convert back corresponding the input form [find in code](../src/TTTEngine/3TEngine.go#L68-L119)
8. 1. Generate all the 8 possible canonical form [find in code](../src/TTTEngine/3TEngine.go#L79)
   2. Take the XOR of the input state's (18 bit of X and O, i.e from right to left (odering from 1), take $10^{th}$ to $27^{th}$ bit) and the same of Canonical best move state.</br> This will give us a binary number with only 1 set bit. [find in code](../src/TTTEngine/3TEngine.go#L112)
   3. The index at which the bit is set is our answer. But this index can be any number between 1 & 18.</br> Take a loop i from 0 to 17 and generate a number 1 << i</br> if this number matches with the XOR result, return index = i+1 (index ordering from 1) 
   4. We, need to convert 1->1, 2->2, 3->3, ... ,9->9, 10->1, 11->2, 12-3, ... ,18-1</br>So, take the modulus of the (index-1)%9+1
   5. But this will give the cell number to play in reverse, so subtract it by 10.</br> i.e 10 - ((index-1)%9+1) </br> $\because$ index = i+1</br> we can return 10 - (i%9+1) [find in code](../src/TTTEngine/3TEngine.go#L104-112)