import {useEffect, useState} from "react";
import {CopyBlock, hybrid} from "react-code-blocks";

function CodePage() {
    const [code, setCode] = useState({})

    useEffect(() => {
        fetch('http://localhost:8000/AoC-2021' + window.location.pathname)
            .then(resp => resp.text())
            .then(codeText => setCode(codeText.toString()))
    }, [window.location.pathname])

    return (
        <div className="text-sm">
            <CopyBlock
                text={code}
                language="python"
                theme={hybrid}
                codeBlock
            />
        </div>
    )
}

export default CodePage