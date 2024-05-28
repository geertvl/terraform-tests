param (
    [string]$DirectoryToAdd
)

if (-not $DirectoryToAdd) {
    Write-Host "Please provide the directory to add to PATH."
    exit 1
}

# Get the current PATH environment variable
$path = [System.Environment]::GetEnvironmentVariable("PATH", "User")

# Split the PATH into an array of directories
$pathArray = $path -split ';'

# Check if the directory is already in PATH
if ($pathArray -contains $DirectoryToAdd) {
    Write-Host "The directory '$DirectoryToAdd' is already in the PATH."
} else {
    # Add the directory to the PATH
    $newPath = "$path;$DirectoryToAdd"
    [System.Environment]::SetEnvironmentVariable("PATH", $newPath, "User")
    Write-Host "The directory '$DirectoryToAdd' has been added to the PATH."
}

# Optional: Verify the update
$updatedPath = [System.Environment]::GetEnvironmentVariable("PATH", "User")
Write-Host "Updated PATH: $updatedPath"
