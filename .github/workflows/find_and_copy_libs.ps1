param (
    [string]$GithubSha,
    [string]$MatrixOs,
    [string]$MatrixBuildType
)

$ErrorActionPreference = "Stop"

$libNamesToFind = @('comsuppw', 'dbghelp', 'msvcrt', 'oldnames', 'ole32', 'uuid', 'oleaut32', 'vcruntime', 'ucrt')
$basePackageDir = "Dawn-Static-$GithubSha-$MatrixOs-$MatrixBuildType"
$targetLibDir = Join-Path -Path $basePackageDir -ChildPath 'lib'

Write-Host "Base Package Dir: $basePackageDir"
Write-Host "Target Lib Dir: $targetLibDir"

# Ensure base package directory and target lib directory exist
New-Item -ItemType Directory -Force -Path $basePackageDir -ErrorAction SilentlyContinue | Out-Null
New-Item -ItemType Directory -Force -Path $targetLibDir -ErrorAction SilentlyContinue | Out-Null

$allLibsFound = $true
Write-Host "Searching for libraries in paths from `$env:LIB (System), `$env:LIBPATH (System), and current directory."

# Combine LIB, LIBPATH, and current directory for search
$searchPaths = ($env:LIB -split ';') + ($env:LIBPATH -split ';') + ($pwd | Select-Object -ExpandProperty Path)
$uniqueSearchPaths = $searchPaths | Sort-Object -Unique | Where-Object { $_ } # Filter out empty paths

foreach ($libBaseName in $libNamesToFind) {
    $foundPath = $null
    # Try .lib first (MSVC standard)
    foreach ($searchDir in $uniqueSearchPaths) {
        $libFilePath = ""
        try {
            # Resolve-Path can error if the base path doesn't exist, so we check dir existence first.
            if (Test-Path $searchDir -PathType Container) {
                $potentialPath = Join-Path -Path $searchDir -ChildPath ($libBaseName + '.lib')
                if (Test-Path $potentialPath -PathType Leaf) {
                    $libFilePath = (Resolve-Path $potentialPath).Path
                }
            }
        } catch {
            # Silently continue if Resolve-Path fails for any reason
        }

        if ($libFilePath) {
            Write-Host ("Found '$($libBaseName).lib' at '$($libFilePath)'")
            Copy-Item -Path $libFilePath -Destination $targetLibDir -Force
            $foundPath = $libFilePath
            break
        }
    }

    # If .lib not found, try .a
    if (-not $foundPath) {
        foreach ($searchDir in $uniqueSearchPaths) {
            $libFilePathA = ""
            try {
                 if (Test-Path $searchDir -PathType Container) {
                    $potentialPathA = Join-Path -Path $searchDir -ChildPath ($libBaseName + '.a')
                    if (Test-Path $potentialPathA -PathType Leaf) {
                        $libFilePathA = (Resolve-Path $potentialPathA).Path
                    }
                }
            } catch {
                 # Silently continue
            }
            if ($libFilePathA) {
                Write-Host ("Found '$($libBaseName).a' at '$($libFilePathA)' (Note: .a is less common for MSVC toolchain)")
                Copy-Item -Path $libFilePathA -Destination $targetLibDir -Force
                $foundPath = $libFilePathA
                break
            }
        }
    }

    if (-not $foundPath) {
        Write-Warning "Could not find library: '$($libBaseName)' (searched for .lib and .a extensions in specified paths)"
        $allLibsFound = $false
    }
}

if ($allLibsFound) {
    Write-Host "All specified libraries were found and copied to '$($targetLibDir)'."
} else {
    Write-Warning "Not all specified libraries were found. Please check warnings above."
    # To make the job fail if a library is not found, uncomment the following line:
    # exit 1
}

Write-Host "Contents of '$($targetLibDir)' after copy operation:"
Get-ChildItem -Path $targetLibDir -ErrorAction SilentlyContinue | ForEach-Object { Write-Host ("  - " + $_.Name) }

# Exit with 0 if all successful, or if $allLibsFound is false and we haven't explicitly exited with 1
if (-not $allLibsFound) {
    # Optionally set a specific exit code if you want to differentiate this "soft" failure
    # For now, it will follow the script's natural exit code (0 if no 'exit 1' was hit)
} 