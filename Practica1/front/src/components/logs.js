import React, { useEffect, useState } from 'react'
import { OperationsTable } from './Tabla'



const baseUrl = 'http://localhost:4000';
export const Reports = () => {

    const [showButton, setshowButton] = useState(true);
    const [operations, setOperations] = useState([])

    useEffect(() => {
        getOperations();
    }, [])

    const getOperations = async() => {
        await fetch(`${baseUrl}/Operaciones`, {
            method: 'GET',
            mode: 'cors',
            headers: {
                "Content-Type": "application/json"
            }
        })
        .then(resp => resp.json())
        .then(data => {
            setOperations(data.reverse())
        }).catch(console.error)
    }



    return (
        <div style={{ width: '70%' }}>
            <br />
            <div
                className='d-flex justify-content-between'
            >
                <h1>Reporte</h1>
                
            </div>
            <br />
            {
                operations.length > 0 ? (
                    <>
                        {(<OperationsTable operations={operations} />)}
                    </>
                ) : (
                    <pre>No hay operaciones(</pre>
                )
            }

        </div>
    )
}