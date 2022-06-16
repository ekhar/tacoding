//using piston api
//it needs the language, version, files at minimum
function compileCode(){
    let language="python"
    let version = "3.10.0"
    let fname = "main"
    let fcontent = "print(5+19)"
    let send = {
        "language":language,
        "version":version,
        "files" : [
            {"name": fname, "content": fcontent }
        ]

    }


    fetch('https://emkc.org/api/v2/piston/execute', {
        method: 'POST',
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(send)
        }
    ).then(response=> response.text()).then(data=>console.log(JSON.parse(data)))

    console.log("hello :)")
}

export const RunButton = () => {
    return (
        <div>
            <button onClick={compileCode}>
                Compile Code
            </button>

        </div>
    )

}

