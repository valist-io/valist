import React, { FunctionComponent, useState, useEffect } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            '& > *': {
                margin: theme.spacing(1),
                width: '45ch',

        },
    },
    }),
);

export const CreateRepoForm:FunctionComponent<any> = ({valist}) => {

    const [account, setAccount] = useState("");
    const [orgName, setOrgName] = useState("")
    const [repoName, setRepoName] = useState("")
    const [repoDescription, setRepoDescription] = useState("")

    const classes = useStyles();

    useEffect(() => {
        if (valist) {
            (async function () {
                try {
                    const accounts = await valist.web3.eth.getAccounts();
                    setAccount(accounts[0]);
                    setOrgName("")
                    setRepoName("")
                    setRepoDescription("")
                } catch (error) {
                    alert(`Failed to load accounts.`);
                    console.log(error);
                }
            })();
        }
    }, [valist]);

    const createRepo = async () => {
        const repoMeta = {
            name: repoName,
            description: repoDescription
        };

        await valist.createRepository(orgName, repoName, repoMeta, account)
    }

    return (
        <div id="org-card">
            <div className="repo-image"></div>
            <Card>
                <form className={classes.root} noValidate autoComplete="off">
                    <TextField onChange={(e) => setOrgName(e.target.value)} id="outlined-basic" label="Organization Name" variant="outlined" value={orgName}/>
                    <br></br>
                    <TextField onChange={(e) => setRepoName(e.target.value)} id="outlined-basic" label="Repo Name" variant="outlined" value={repoName} />
                    <br></br>
                    <TextField onChange={(e) => setRepoDescription(e.target.value)} id="outlined-basic" label="Description" variant="outlined" value={repoDescription} />
                    <br></br>
                    <Button onClick={createRepo} value="Submit" variant="contained" color="primary">Create</Button>
                    </form>
            </Card>
        </div>
    );
}
