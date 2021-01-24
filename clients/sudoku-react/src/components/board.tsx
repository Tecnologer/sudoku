import React, { ChangeEvent } from "react";
import "../App.css";

interface ICoordinate {
  x: number;
  y: number;
  val: string;
  isLocked: boolean;
}
export interface BoardProps {
  content: string[][];
}

export interface BoardState {
  content: ICoordinate[][];
}

class Board extends React.Component<BoardProps, BoardState> {
  componentWillReceiveProps(nextProps: BoardProps) {
    this.setState({
      content: this.parseContentToCoordinates(nextProps.content),
    });
  }

  state = { content: Array<ICoordinate[]>() };
  coordinates = Array<ICoordinate[]>();

  render() {
    const { content } = this.props;
    this.coordinates = this.parseContentToCoordinates(content);
    return (
      <div className="board">{this.getCoordinateFields(this.coordinates)}</div>
    );
  }

  componentDidMount() {
    this.setState({ content: this.coordinates });
  }

  getCoordinateFields = (content: ICoordinate[][]) => {
    return content.map((row, x) => {
      return (
        <div id={"row_" + x}>
          {row.map((col, y) => {
            return (
              <div
                id={"col_" + x + "_" + y}
                className={this.getCoordinateClass(x, y)}
              >
                {this.formatCell(col)}
              </div>
            );
          })}
        </div>
      );
    });
  };

  formatCell = (coor: ICoordinate) => {
    if (coor.isLocked) {
      return <span id={"cell_" + coor.x + "_" + coor.y}>{coor.val}</span>;
    }

    return (
      <input
        id={coor.x + "," + coor.y}
        type="text"
        className="coordinate-editable "
        value={coor.val}
        onChange={this.handleChange}
      />
    );
  };

  getCoordinateClass = (x: number, y: number) => {
    let c: string = "coordinate";

    if (y === 0) {
      c += " coordinate-start";
    } else if (y === 8) {
      c += " coordinate-end";
    } else {
      c += " coordinate-middle";
    }

    if (x === 8) {
      c += " coordinate-end-row";
    }

    if ((x + 1) % 3 === 0) {
      c += " coordinate-bolder-row";
    }
    // console.log(x, y, x % 3, y % 3);
    if ((y + 1) % 3 === 0) {
      c += " coordinate-bolder-column";
    }

    return c;
  };

  handleChange = (event: ChangeEvent<{ value: string; id: string }>) => {
    let id = event.target.id.split(",");
    console.log(id);
    let x: number = Number(id[0]);
    let y: number = Number(id[1]);
    let content: ICoordinate[][] = this.state.content;
    console.log(content);
    content[x][y].val = event.target.value;
    console.log(content);
    this.setState({ content: content });
  };

  parseContentToCoordinates = (content: string[][]) => {
    let coordinates: ICoordinate[][] = [];
    content.forEach((row, x) => {
      coordinates[x] = [];
      row.forEach((col, y) => {
        coordinates[x][y] = {
          x: x,
          y: y,
          val: col,
          isLocked: col !== "" && col !== "0",
        };
      });
    });

    return coordinates;
  };
}

export default Board;
