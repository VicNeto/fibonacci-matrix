import React, {Component} from "react";
import axios from "axios";

export class Spiral extends Component {

    constructor() {
        super();
        this.state = {
            rows: 5,
            columns: 3,
            spiral: [],
         };

        this.requestSpiral = this.requestSpiral.bind(this);
        this.handleChangeColumnInput = this.handleChangeColumnInput.bind(this);
        this.handleChangeRowInput = this.handleChangeRowInput.bind(this);
    }

    requestSpiral() {
        const rows = this.state.rows;
        const cols = this.state.columns;
        axios.get(`${process.env.REACT_APP_API_URL}/spiral?rows=${rows}&cols=${cols}`)
            .then(resp => this.setState({
                spiral: resp.data,
            }))
            .catch(err => console.log(err));
    }

    renderRow(row, i) {
        return (
            <tr key={i} className="bg-emerald-200">
                {row?.map((number, j) => 
                    <td key={j} className="w-1/12 border border-green-600">
                        <p className="px-1">{number}</p>
                    </td>)}
            </tr>
        )
    }

    handleChangeRowInput(event) {
        this.setState({ rows: event.target.value });
    }

    handleChangeColumnInput(event) {
        this.setState({ columns: event.target.value });
    }

    sizeForm() {
        return (
            <div>
                <label className="px-2">Number of rows</label>
                <input 
                    className="border-green-800 border" 
                    type="text" 
                    value={this.state.rows} 
                    onChange={this.handleChangeRowInput}>
                </input>
                <label className="px-2">Number of columns</label>
                <input 
                    className="border-green-800 border" 
                    type="text" 
                    value={this.state.columns} 
                    onChange={this.handleChangeColumnInput}>
                </input>
                <button className="border border-green-800 ml-2" onClick={this.requestSpiral}>Calculate</button>
            </div>
        )
    }

    render() {
        const form = this.sizeForm();
        const spiral = 
                <tbody>
                    {this.state.spiral?.map((row, i) =>
                        this.renderRow(row, i)
                    )}
                </tbody>
        return (
            <div>    
                {form}
                <br/>
                <table className="table-auto border-collapse border border-green-800">
                    {spiral}
                </table>
            </div>
        );
    }
}
