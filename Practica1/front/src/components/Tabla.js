import React from 'react'
import './Tabla.css';
export const OperationsTable = ({ operations }) => {


    return (
        
            <table className='table'  border = "1">
                
                <thead className='header'>
                    <tr>
                        <th>No.</th>
                        <th>Numero1</th>
                        <th>Numero2</th>
                        <th>Operador</th>
                        <th>Resultado</th>
                        <th>Fecha</th>
                    </tr>
                </thead>
                <tbody className='cuerpo'>
                    {
                        operations.map( ({Numero1, Numero2, Operador, Resultado, Fecha}, i) => {
                            return (
                                <tr>
                                    <td>{i + 1}</td>
                                    <td>{Numero1}</td>
                                    <td>{Numero2}</td>
                                    <td>{Operador}</td>
                                    <td>{Resultado}</td>
                                    <td>{Fecha}</td>
                                </tr>
                            )
                        })
                    }
                </tbody>
            </table>
        
    )
}