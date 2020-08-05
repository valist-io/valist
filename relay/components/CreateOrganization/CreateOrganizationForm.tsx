import React, { FunctionComponent, useState, useEffect } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

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

export const CreateOrganizationForm:FunctionComponent<any> = ({valist}) => {

    const [account, setAccount] = useState("");

    const [orgShortName, setOrgShortName] = useState("")
    const [orgFullName, setOrgFullName] = useState("")
    const [orgDescription, setOrgDescription] = useState("")

    const classes = useStyles();

    useEffect(() => {
        if (valist) {
            (async function () {
                try {
                    const accounts = await valist.web3.eth.getAccounts();
                    setAccount(accounts[0]);
                } catch (error) {
                    alert(`Failed to load accounts.`);
                    console.log(error);
                }
            })();
        }
    }, [valist]);

    const createOrganization = async () => {
        const meta = {
            name: orgFullName,
            description: orgDescription
        };

        await valist.createOrganization(orgShortName, JSON.stringify(meta), account);
    }

    return (
        <form className={classes.root} noValidate autoComplete="off">
            <TextField onChange={(e) => setOrgShortName(e.target.value)} id="outlined-basic" label="Organization Short Name" variant="outlined" />
            <br></br>
            <TextField onChange={(e) => setOrgFullName(e.target.value)} id="outlined-basic" label="Organization Full Name" variant="outlined" />
            <br></br>
            <TextField onChange={(e) => setOrgDescription(e.target.value)} id="outlined-basic" label="Organization Name" variant="outlined" />
            <br></br>
            <Button onClick={createOrganization} value="Submit" variant="contained" color="primary">Create</Button>
        </form>
    );

}
